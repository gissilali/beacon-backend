package datastore

import (
	"beacon.silali.com/internal/api/config"
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func OpenConnection(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)

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
