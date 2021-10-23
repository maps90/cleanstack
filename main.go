package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/maps90/cleanstack/internal/presenter"

	"github.com/labstack/echo/middleware"
	"github.com/maps90/cleanstack/pkg/datasources/mysql"
	"github.com/maps90/cleanstack/pkg/transport/httpx"
	"github.com/spf13/viper"
)

var (
	cfgFile *string
	cfgPort *string
)

func init() {
	cfgFile = flag.String("config", "config.yaml", "config file (default is {WORKSPACE}/config.yaml)")
	cfgPort = flag.String("port", "3000", "default port is :3000")
	flag.Parse()

	// Use config file from the flag.
	viper.SetConfigFile(*cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// mysql initialize
	mysql.Init()
}

func main() {
	// instantiate new httpx package
	instance := httpx.New()
	instance.HideBanner = true
	instance.SetPort(":" + *cfgPort)

	instance.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.Gzip(),
	)

	// override http error handler with exception package
	instance.Echo.HTTPErrorHandler = httpx.ErrorHandler

	// serve the http server
	instance.Serve()
}
