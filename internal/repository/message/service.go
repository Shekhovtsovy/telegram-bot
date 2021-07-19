package message

import "database/sql"

// Repository is an interface which provides methods for message repository work
type Repository interface {
	InsertMessage() error
}

type repository struct {
	db *sql.DB
}

// Insert message to database
func (r *repository) InsertMessage() error {
	println("Insert Message To Database")
	return nil
}

// NewRepository return a new message repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
