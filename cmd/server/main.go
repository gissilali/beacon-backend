package main

import (
	"beacon.silali.com/internal/api"
	"beacon.silali.com/internal/api/config"
	"beacon.silali.com/internal/api/core"
	"beacon.silali.com/internal/api/data"
	"beacon.silali.com/internal/api/datastore"
	"database/sql"
	"flag"
	"log"
	"os"
)

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 8002, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|prod)")
	flag.StringVar(&cfg.DB.DSN, "database-dsn", "postgres://beacon:password@db:5432/beacon?sslmode=disable", "Postgres DSN")

	logger := log.New(os.Stdout, "BEACON~~~", log.Ldate|log.Ltime)

	db, err := datastore.OpenConnection(cfg)
	if err != nil {
		logger.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Printf("Failed to close the database connection %s", err)
		}
	}(db)

	appCtx := core.New(&cfg, logger, data.NewModel(db), "1")

	serverErr := api.StartServer(&cfg, appCtx)
	if serverErr != nil {
		logger.Println(serverErr)
	}
}
