package message

import (
	"bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// Service is an interface which provides methods for message service work
type Service interface {
	SaveMessage(message *tgbotapi.Message) error
}

type repository interface {
	InsertMessage(message *tgbotapi.Message) error
}

type service struct {
	repository repository
	log        logger.Log
}

// Save message
func (s *service) SaveMessage(msg *tgbotapi.Message) error {
	if err := s.repository.InsertMessage(msg); err != nil {
		s.log.Error("message save error", zap.Int("id", msg.MessageID))
		return err
	}
	s.log.Info("message saved", zap.Int("id", msg.MessageID))
	return nil
}

// NewService return a new message service
func NewService(r repository) Service {
	l := logger.NewLogger("messageService")
	return &service{
		repository: r,
		log:        l,
	}
}
