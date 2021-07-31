package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Service is an interface which provides methods for message service work
type Service interface {
	SaveMessage(msg *tgbotapi.Message) (*tgbotapi.Message, error)
}

type repository interface {
	InsertMessage(message *tgbotapi.Message) error
}

type service struct {
	repository repository
}

// Save message
func (s *service) SaveMessage(msg *tgbotapi.Message) (*tgbotapi.Message, error) {
	if err := s.repository.InsertMessage(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// NewService return a new message service
func NewService(r repository) Service {
	return &service{
		repository: r,
	}
}
