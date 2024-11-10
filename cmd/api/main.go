package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config  config
	logger  *log.Logger
	version string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8002, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config:  cfg,
		logger:  logger,
		version: version,
	}

	e := echo.New()
	e.GET("/v1/healthcheck", app.healthcheckHandler)
	e.GET("/v1/servers", app.healthcheckHandler)
	logger.Fatal(e.Start(fmt.Sprintf(":%d", app.config.port)))
}
