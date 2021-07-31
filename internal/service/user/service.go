package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Service is an interface which provides methods for user service work
type Service interface {
	SaveUserIfNew(user *tgbotapi.User) (*tgbotapi.User, error)
}

type repository interface {
	InsertUser(user *tgbotapi.User) error
	SelectUser(id int) (tgbotapi.User, error)
}

type service struct {
	repository repository
}

// Save user
func (s *service) SaveUserIfNew(user *tgbotapi.User) (*tgbotapi.User, error) {
	_, err := s.repository.SelectUser(user.ID)
	if err != nil {
		if err := s.repository.InsertUser(user); err != nil {
			return nil, err
		}
	}
	return user, nil
}

// NewService return a new message service
func NewService(r repository) Service {
	return &service{
		repository: r,
	}
}
