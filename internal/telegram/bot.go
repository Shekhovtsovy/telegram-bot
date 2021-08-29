package telegram

import (
	"bot/internal/config"
	"bot/internal/logger"
	"bot/internal/model"
	"bot/internal/system/metrics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"strings"
)

// Bot is an interface which provides methods for bot working
type Bot interface {
	Listen()
}

type userRep interface {
	GetOne(userId int) (model.User, error)
	AddOne(user *tgbotapi.User) error
}

type msgRep interface {
	AddOne(msg *tgbotapi.Message) (*tgbotapi.Message, error)
	GetAll(userId int, chatId int) ([]model.Message, error)
	DeleteOne(msgId int, chatId int) error
}

type bot struct {
	api     *tgbotapi.BotAPI
	cfg     config.Telegram
	stat    metrics.TelegramStat
	userRep userRep
	msgRep  msgRep
	log     logger.Log
}

// Listen income data
func (b *bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		b.stat.IncTelegramListeningErrors()
		b.log.Error("telegram listening error", zap.String("details", err.Error()))
	}
	for update := range updates {
		if update.CallbackQuery != nil {
			if err := b.processCallback(update.CallbackQuery); err != nil {
				b.log.Error("processing callback error", zap.String("details", err.Error()))
			}
		}
		if update.Message == nil {
			continue // ignore any non-Message Updates
		}
		b.stat.IncReceivedMessages()
		b.log.Info("got message from telegram",
			zap.String("text", update.Message.Text),
			zap.Int("messageId", update.Message.From.ID),
			zap.Int("userId", update.Message.From.ID))
		b.saveNewUser(update.Message.From)
		b.saveIncomeMessage(update.Message)
		if err := b.processMessage(update.Message); err != nil {
			b.log.Error("processing message error", zap.String("details", err.Error()))
		}
	}
}

// saveNewUser saves user if he does not exist in database
func (b *bot) saveNewUser(user *tgbotapi.User) {
	if _, err := b.userRep.GetOne(user.ID); err != nil {
		b.stat.IncSavingUserErrors()
		b.log.Error("saving user error", zap.String("details", err.Error()))
	} else {
		if err := b.userRep.AddOne(user); err != nil {
			b.stat.IncSavingUserErrors()
			b.log.Error("saving user error", zap.String("details", err.Error()))
		} else {
			b.stat.IncNewUsers()
			b.log.Info("user saved", zap.Int("userId", user.ID))
		}
	}
}

// saveIncomeMessage saves income message
func (b *bot) saveIncomeMessage(msg *tgbotapi.Message) {
	if msg, err := b.msgRep.AddOne(msg); err != nil {
		b.stat.IncSavingMessageErrors()
		b.log.Error("saving message error", zap.String("details", err.Error()))
	} else {
		b.log.Info("message saved", zap.Int("messageId", msg.MessageID))
	}
}

// processMessage processes an income message
func (b *bot) processMessage(msg *tgbotapi.Message) error {
	switch msg.Text {
	case commandStart:
		if err := b.handleCommandStart(msg); err != nil {
			return err
		}
	case commandAbout:
		if err := b.handleCommandAbout(msg); err != nil {
			return err
		}
	default:
		if err := b.handleUnknownCommand(msg); err != nil {
			return err
		}
	}
	return nil
}

// processCallback processes a callback
func (b *bot) processCallback(c *tgbotapi.CallbackQuery) error {
	s := strings.Split(c.Data, "|")
	switch s[1] { // 0 - data, 1 - callback key, 2 - callback sub key, ...
	case callbackStart:
		if err := b.handleCommandStart(c.Message); err != nil {
			return err
		}
	case callbackAbout:
		if err := b.handleCommandAbout(c.Message); err != nil {
			return err
		}
	default:
		if err := b.handleUnknownCommand(c.Message); err != nil {
			return err
		}
	}
	return nil
}

// NewBot return new Telegram Bot
func NewBot(cfg config.Config, mr msgRep, ur userRep) Bot {
	b, _ := tgbotapi.NewBotAPI(cfg.Telegram.ApiToken)
	b.Debug = cfg.Telegram.IsLog
	return &bot{
		api:     b,
		cfg:     cfg.Telegram,
		stat:    metrics.NewTelegramStat(),
		userRep: ur,
		msgRep:  mr,
		log:     logger.NewLogger("bot"),
	}
}
