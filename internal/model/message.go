package model

import (
	"time"
)

type Message struct {
	Id        uint
	UserId    uint
	ChatId    uint
	Text      string
	CreatedAt *time.Time
}
