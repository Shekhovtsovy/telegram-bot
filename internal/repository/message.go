package repository

import (
	"bot/internal/model"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const messageTableName = "messages"

// Message is an interface which provides methods for message repository work
type Message interface {
	SaveOne(msg *tgbotapi.Message) (*tgbotapi.Message, error)
	GetAll(userId int, chatId int) ([]model.Message, error)
	DeleteOne(msgId int, chatId int) error
}

type message struct {
	db *sql.DB
}

// SaveOne inserts a message to database
func (r *message) SaveOne(msg *tgbotapi.Message) (*tgbotapi.Message, error) {
	createdAt := time.Unix(int64(msg.Date), 0).Format("2006-01-02 15:04:05")
	query := fmt.Sprintf(`INSERT INTO "%s"("id", "created_at", "user_id", "chat_id", "text") values($1, $2, $3, $4, $5)`, messageTableName)
	if _, err := r.db.Exec(query, msg.MessageID, createdAt, msg.From.ID, msg.Chat.ID, msg.Text); err != nil {
		return nil, err
	}
	return msg, nil
}

// GetAll select and returns all messages from database by user ID and chat ID
func (r *message) GetAll(userId int, chatId int) ([]model.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE user_id=$1 AND chat_id=$2`, messageTableName)
	rows, err := r.db.Query(query, userId, chatId)
	var msgs []model.Message
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		msg := model.Message{}
		if err := rows.Scan(&msg.Id, &msg.CreatedAt, &msg.UserId, &msg.ChatId, &msg.Text); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

// DeleteOne deletes a message from database by message ID and chat ID
func (r *message) DeleteOne(msgId int, chatId int) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id=$1 AND chat_id=$2`, messageTableName)
	if _, err := r.db.Exec(query, msgId, chatId); err != nil {
		return err
	}
	return nil
}

// NewMessage returns new message repository
func NewMessage(db *sql.DB) Message {
	return &message{
		db: db,
	}
}
