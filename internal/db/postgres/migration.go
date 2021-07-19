package postgres

import (
	"bot/internal/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigration() (*migrate.Migrate, error) {
	cfg := config.GetConfig()
	m, err := migrate.New(
		"file://internal/db/postgres/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=enable", cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name))
	return m, err
}

// Migrate
func Up() error {
	migration, err := getMigration()
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		switch fmt.Sprintf("%s", err) {
		default:
			return err
		case "no change":
			return nil
		}
	}

	return nil
}

// Rollback migration
func Down() error {
	migration, err := getMigration()
	if err != nil {
		return err
	}

	if err := migration.Down(); err != nil {
		switch fmt.Sprintf("%s", err) {
		default:
			return err
		case "no change":
			return nil
		}
	}

	return nil
}
