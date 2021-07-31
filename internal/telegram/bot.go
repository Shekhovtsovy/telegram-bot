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
		b.stat.IncReceivedMessages()
		b.log.Info("got message from telegram",
			zap.String("text", update.Message.Text),
			zap.Int("message_id", update.Message.From.ID),
			zap.Int("user_id", update.Message.From.ID))
		user, userErr := b.userService.SaveUserIfNew(update.Message.From)
		if userErr == nil {
			b.log.Info("user saved", zap.Int("user_id", user.ID))
		} else {
			b.log.Error("user save error")
		}
		msg, msgErr := b.messageService.SaveMessage(update.Message)
		if msgErr == nil {
			b.log.Info("message saved", zap.Int("message_id", msg.MessageID))
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
	switch msg.Text {
	case commandStartRequest:
		if err := b.handleCommandStart(msg); err != nil {
			return err
		}
	case commandStopRequest:
		if err := b.handleCommandStop(msg); err != nil {
			return err
		}
	case commandAboutRequest:
		if err := b.handleCommandAbout(msg); err != nil {
			return err
		}
	default:
		answer := tgbotapi.NewMessage(msg.Chat.ID, "❌ command not found")
		_, err := b.api.Send(answer)
		if err != nil {
			return err
		}
		b.log.Info("send message from bot")
	}
	return nil
}

// NewBot return a new Telegram Bot
func NewBot(cfg config.Config, ms messageService, us userService) Bot {
	b, _ := tgbotapi.NewBotAPI(cfg.Telegram.ApiToken)
	b.Debug = cfg.Telegram.IsLog
	l := logger.NewLogger("bot")
	return &bot{
		api:            b,
		name:           cfg.Telegram.BotName,
		stat:           metrics.NewTelegramStat(),
		messageService: ms,
		userService:    us,
		log:            l,
	}
}
