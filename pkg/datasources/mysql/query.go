package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	// SQL struct
	SQL struct {
		DB      *sqlx.DB
		logMode bool
		context context.Context
	}
)

// WithContext add context to sql
func (sql *SQL) WithContext(ctx context.Context) *SQL {
	return &SQL{
		DB:      sql.DB,
		logMode: sql.logMode,
		context: ctx,
	}
}

// Query queries the database and returns an *sql.Rows.
func (sql *SQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.QueryContext(sql.context, query, args...)
	}

	return sql.DB.Query(query, args...)
}

// QueryRow queries the database and returns an *sqlx.Row.
func (sql *SQL) QueryRow(query string, args ...interface{}) *sql.Row {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.QueryRowContext(sql.context, query, args...)
	}
	return sql.DB.QueryRow(query, args...)
}

// Queryx queries the databas	e and returns an *sqlx.Rows.
func (sql *SQL) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.QueryxContext(sql.context, query, args...)
	}

	return sql.DB.Queryx(query, args...)
}

// QueryRowx queries the database and returns an *sqlx.Row.
func (sql *SQL) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.QueryRowxContext(sql.context, query, args...)
	}
	return sql.DB.QueryRowx(query, args...)
}

// Exec using master sql
func (sql *SQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.ExecContext(sql.context, query, args...)
	}
	return sql.DB.Exec(query, args...)
}

// Select using slave sql.
func (sql *SQL) Select(dest interface{}, query string, args ...interface{}) error {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.SelectContext(sql.context, dest, query, args...)
	}

	return sql.DB.Select(dest, query, args...)
}

// Get using slave sql.
func (sql *SQL) Get(dest interface{}, query string, args ...interface{}) error {
	Log(sql.logMode, query, args...)
	if sql.context != nil {
		return sql.DB.GetContext(sql.context, dest, query, args...)
	}

	return sql.DB.Get(dest, query, args...)
}

// MustBegin starts a transaction, and panics on error.
func (sql *SQL) MustBegin() *Tx {
	tx, err := sql.DB.Beginx()
	if err != nil {
		panic(err)
	}

	return &Tx{Tx: tx, logMode: sql.logMode, context: sql.context}
}

// Begin begins a transaction and returns an *Tx instead of an *sql.Tx.
func (sql *SQL) Begin() (*Tx, error) {
	tx, err := sql.DB.Beginx()
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx, logMode: sql.logMode, context: sql.context}, nil
}
