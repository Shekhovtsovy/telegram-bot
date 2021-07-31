package message

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const tableName = "messages"

// Repository is an interface which provides methods for message repository work
type Repository interface {
	InsertMessage(msg *tgbotapi.Message) error
}

type repository struct {
	db *sql.DB
}

// Insert message to database
func (r *repository) InsertMessage(msg *tgbotapi.Message) error {
	createdAt := time.Unix(int64(msg.Date), 0).Format("2006-01-02 15:04:05")
	if _, err := r.db.Exec(fmt.Sprintf(`INSERT INTO "%s"("id", "created_at", "user_id", "chat_id", "text") values($1, $2, $3, $4, $5)`,
		tableName), msg.MessageID, createdAt, msg.From.ID, msg.Chat.ID, msg.Text); err != nil {
		return err
	}
	return nil
}

// NewRepository return a new message repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
