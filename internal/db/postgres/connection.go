package postgres

import (
	"bot/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

// GetDb returns database connection
func GetDb(cfg config.Config) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name))
	return db, err
}
