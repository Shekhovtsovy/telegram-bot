package message

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Repository is an interface which provides methods for message repository work
type Repository interface {
	InsertMessage(message *tgbotapi.Message) error
}

type repository struct {
	db *sql.DB
}

// Insert message to database
func (r *repository) InsertMessage(message *tgbotapi.Message) error {
	return nil
}

// NewRepository return a new message repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
