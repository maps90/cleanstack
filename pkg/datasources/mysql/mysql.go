// Package mysql storage is designed to give lazy load singleton access to mysql connections
// it doesn't provide any cluster nor balancing support, assuming it is handled
// in lower level infra, i.e. proxy, cluster etc.
package mysql

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // use by sqlx
	"github.com/jmoiron/sqlx"
	"github.com/maps90/cleanstack/pkg/utils/defaults"
	"github.com/spf13/viper"
)

const (
	readDBKey  = "read"
	writeDBKey = "write"
)

var dbConn *Conn

// Conn struct
type Conn struct {
	sqlx map[string]*SQL
	mux  sync.Mutex
}

// Config struct
type Config struct {
	User        string
	Password    string
	Address     string
	DB          string
	LogMode     bool
	MaxOpen     int
	MaxIdle     int
	MaxLifetime int
}

func Init() error {
	if dbConn == nil {
		dbConn = &Conn{
			sqlx: make(map[string]*SQL),
		}
	}

	return nil
}

func New() error {
	if dbConn == nil {
		dbConn = &Conn{
			sqlx: make(map[string]*SQL),
		}
	}

	return nil
}

// get connection in thread-safe fashion
func (db *Conn) get(id string) *SQL {
	db.mux.Lock()
	defer db.mux.Unlock()
	if conn, ok := db.sqlx[id]; ok {
		return conn
	}
	return nil
}

// set connection in thread-safe fashion
func (db *Conn) set(id string, sqlx *SQL) {
	db.mux.Lock()
	db.sqlx[id] = sqlx
	db.mux.Unlock()
}

// Read retrieve MySQL established connection client (sqlx) and panic if error
func Read() *SQL {
	if d, err := dbConn.connect(readDBKey); err != nil {
		panic(err)
	} else {
		return d
	}
}

// Write retrieve MySQL established connection client (sqlx) and panic if error
func Write() *SQL {
	if d, err := dbConn.connect(writeDBKey); err != nil {
		panic(err)
	} else {
		return d
	}
}

func (db *Conn) connect(id string) (*SQL, error) {
	mysqlConfig := viper.GetStringMap("mysql")
	if _, ok := mysqlConfig[id]; !ok {
		return nil, fmt.Errorf("mysql configuration for [%s] does not exists", id)
	}

	// if previously established, reuse and ping
	if con := db.get(id); con != nil {
		return con, nil
	}

	con, err := db.openConnection(id)
	if err != nil {
		return nil, err
	}

	// otherwise establish new connection through centralized component config
	db.set(id, con)
	return db.get(id), nil
}

// Shutdown disconnecting all established mysql client connection
func Shutdown() (err error) {
	if dbConn == nil {
		return nil
	}
	for _, c := range dbConn.sqlx {
		err = c.DB.Close()
	}
	return err
}

func (db *Conn) openConnection(id string) (*SQL, error) {
	opt := setupConfig(id)
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		opt.User,
		opt.Password,
		opt.Address,
		opt.DB,
	)

	con, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	con.SetConnMaxLifetime(time.Duration(opt.MaxLifetime))
	con.SetMaxOpenConns(opt.MaxOpen)
	con.SetMaxIdleConns(opt.MaxIdle)

	db.set(id, &SQL{DB: con, logMode: opt.LogMode})
	return db.get(id), nil
}

func setupConfig(id string) *Config {
	option := &Config{
		User:     viper.GetString(getKey(id, "user")),
		Password: viper.GetString(getKey(id, "password")),
		Address:  viper.GetString(getKey(id, "address")),
		DB:       viper.GetString(getKey(id, "db")),
	}

	option.MaxLifetime = defaults.Int(viper.GetInt("mysql.max_lifetime"), 30)
	option.MaxOpen = defaults.Int(viper.GetInt("mysql.max_open"), 30)
	option.LogMode = viper.GetBool("mysql.log")

	return option
}

func getKey(id, types string) string {
	return fmt.Sprintf("mysql.%s.%s", id, types)
}
