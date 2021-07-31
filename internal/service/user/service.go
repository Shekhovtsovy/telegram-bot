package user

import (
	"bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Service is an interface which provides methods for user service work
type Service interface {
	SaveUserIfNew(user *tgbotapi.User) error
}

type repository interface {
	InsertUser(user *tgbotapi.User) error
	SelectUser(id int) (tgbotapi.User, error)
}

type service struct {
	repository repository
	log        logger.Log
}

// Save user
func (s *service) SaveUserIfNew(user *tgbotapi.User) error {
	_, err := s.repository.SelectUser(user.ID)
	if err != nil {
		if err := s.repository.InsertUser(user); err != nil {
			s.log.Error("user save error")
			return err
		}
	}
	s.log.Info("user saved")
	return nil
}

// NewService return a new message service
func NewService(r repository) Service {
	l := logger.NewLogger("user")
	return &service{
		repository: r,
		log:        l,
	}
}
