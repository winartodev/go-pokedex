package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewDatabase is function to make connection to database
func NewDatabase(cfg Config) (db *sql.DB, err error) {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	db, err = sql.Open(cfg.Database.Connection, fmt.Sprint(dbConfig, cfg.Database.Database))
	if err != nil {
		return db, err
	}

	return db, err
}
