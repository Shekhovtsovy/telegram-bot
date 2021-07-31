package telegram

import (
	"bot/internal/config"
	"bot/internal/logger"
	"bot/internal/system/metrics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// Bot is an interface which provides methods for bot working
type Bot interface {
	Listen()
}

type messageService interface {
	SaveMessage(msg *tgbotapi.Message) (*tgbotapi.Message, error)
}

type userService interface {
	SaveUserIfNew(user *tgbotapi.User) (*tgbotapi.User, error)
}

type bot struct {
	api            *tgbotapi.BotAPI
	name           string
	stat           metrics.TelegramStat
	messageService messageService
	userService    userService
	log            logger.Log
}

// Listen messages
func (b *bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		b.log.Info("got message from telegram")
		b.stat.IncReceivedMessages()
		user, userErr := b.userService.SaveUserIfNew(update.Message.From)
		if userErr == nil {
			b.log.Info("user saved", zap.Int("id", user.ID))
		} else {
			b.log.Error("user save error")
		}
		msg, msgErr := b.messageService.SaveMessage(update.Message)
		if msgErr == nil {
			b.log.Info("message saved", zap.Int("id", msg.MessageID))
		} else {
			b.log.Error("message save error")
		}
		if err := b.processMessage(update.Message); err != nil {
			b.log.Error("processing message error")
		}
	}
}

// Process message
func (b *bot) processMessage(msg *tgbotapi.Message) error {
	if msg.Text == commandAboutRequest {
		answer := tgbotapi.NewMessage(msg.Chat.ID, commandAboutResponse)
		_, err := b.api.Send(answer)
		if err != nil {
			return err
		}
		b.log.Info("send message from bot", zap.String("command", commandAboutRequest))
	}
	return nil
}

// NewBot return a new Telegram Bot
func NewBot(cfg config.Config, ms messageService, us userService) Bot {
	b, _ := tgbotapi.NewBotAPI(cfg.Telegram.ApiToken)
	b.Debug = cfg.Telegram.IsLog
	l := logger.NewLogger("telegram")
	return &bot{
		api:            b,
		name:           cfg.Telegram.BotName,
		stat:           metrics.NewTelegramStat(),
		messageService: ms,
		userService:    us,
		log:            l,
	}
}
