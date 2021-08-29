package model

import (
	"time"
)

type User struct {
	Id           uint
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
	IsBot        bool
	CreatedAt    *time.Time
}
