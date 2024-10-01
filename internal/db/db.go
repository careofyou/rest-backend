package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
    db, err := sql.Open("postgres", addr)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(maxOpenConns)
    db.SetMaxIdleConns(maxIdleConns)

    duration, err := time.ParseDuration(maxIdleTime)
    if err != nil {
        return nil, err 
    }

    db.SetConnMaxIdleTime(duration)
    
    ctx, cacnel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cacnel()

    if err = db.PingContext(ctx); err != nil {
        return nil, err
    }

    return db, nil
}