package repository

import (
	"bot/internal/model"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const userTableName = "users"

// User is an interface which provides methods for user repository work
type User interface {
	GetOne(userId int) (model.User, error)
	SaveOne(user *tgbotapi.User) error
}

type user struct {
	db *sql.DB
}

// GetOne select and returns the user from database by user ID
func (r *user) GetOne(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id=$1`, userTableName)
	err := r.db.QueryRow(query, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.LanguageCode, &user.IsBot, &user.CreatedAt)
	return user, err
}

// SaveOne inserts a message to database
func (r *user) SaveOne(user *tgbotapi.User) error {
	query := fmt.Sprintf(`INSERT INTO "%s"("id", "first_name", "last_name", "user_name", "language_code", "is_bot", "created_at") values($1, $2, $3, $4, $5, $6, $7)`, userTableName)
	if _, err := r.db.Exec(query, user.ID, user.FirstName, user.LastName, user.UserName, user.LanguageCode, user.IsBot, time.Now()); err != nil {
		return err
	}
	return nil
}

// NewUser returns new user repository
func NewUser(db *sql.DB) User {
	return &user{
		db: db,
	}
}
