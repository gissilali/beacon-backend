package main

import (
	"beacon.silali.com/internal/data"
	"context"
	"database/sql"
	"errors"
	"flag"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config  config
	logger  *log.Logger
	version string
	models  data.Models
}

func (app *application) currentUser(c echo.Context) (*data.User, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}

	claims, claimsIsOk := token.Claims.(jwt.MapClaims)
	if !claimsIsOk {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}

	user := &data.User{}
	if err := user.FromMapClaims(claims); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8002, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|prod)")
	flag.StringVar(&cfg.db.dsn, "database-dsn", "postgres://beacon:password@db:5432/beacon?sslmode=disable", "Postgres DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDBConnection(cfg)

	if err != nil {
		logger.Fatalln(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Printf("Failed to close the database connection %s", err)
		}
	}(db)

	logger.Println("Database connection pool established.")
	models := data.NewModel(db)
	app := &application{
		config:  cfg,
		logger:  logger,
		version: version,
		models:  models,
	}

	app.routes()
}

func openDBConnection(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil

}
