package repository

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	mu sync.Mutex
)

func ConfigureDb(connectionString string) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	if db == nil {
		var err error

		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			return nil, fmt.Errorf("can't open mysql connection pool: %w", err)
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}

	return db, nil
}
