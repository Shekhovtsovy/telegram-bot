package user

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const tableName = "users"

// Repository is an interface which provides methods for user repository work
type Repository interface {
	InsertUser(user *tgbotapi.User) error
	SelectUser(id int) (tgbotapi.User, error)
}

type repository struct {
	db *sql.DB
}

// Select user from database
func (r *repository) SelectUser(id int) (tgbotapi.User, error) {
	var user tgbotapi.User
	var createdAt string
	err := r.db.QueryRow(fmt.Sprintf(`SELECT * FROM "%s" WHERE id=$1`, tableName), id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.LanguageCode, &user.IsBot, &createdAt)
	return user, err
}

// Insert user to database
func (r *repository) InsertUser(user *tgbotapi.User) error {
	if _, err := r.db.Exec(fmt.Sprintf(`INSERT INTO "%s"("id", "first_name", "last_name", "user_name", "language_code", "is_bot") values($1, $2, $3, $4, $5, $6)`,
		tableName), user.ID, user.FirstName, user.LastName, user.UserName, user.LanguageCode, user.IsBot); err != nil {
		return err
	}
	return nil
}

// NewRepository return a new user repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
