package message

import "database/sql"

// Repository is an interface which provides methods for user repository work
type Repository interface {
	InsertUser() error
}

type repository struct {
	db *sql.DB
}

// Insert message to database
func (r *repository) InsertUser() error {
	println("Insert User To Database")
	return nil
}

// NewRepository return a new message repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
