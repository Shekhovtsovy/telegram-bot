package telegram

import (
	"bot/internal/config"
	"bot/internal/system/metrics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot is an interface which provides methods for bot working
type Bot interface {
	Listen()
}

type messageService interface {
	SaveMessage(message *tgbotapi.Message) error
}

type userService interface {
	SaveUserIfNew(user *tgbotapi.User) error
}

type bot struct {
	api            *tgbotapi.BotAPI
	name           string
	stat           metrics.TelegramStat
	messageService messageService
	userService    userService
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
		if err := b.messageService.SaveMessage(update.Message); err != nil {
		}
		if err := b.userService.SaveUserIfNew(update.Message.From); err != nil {
		}
	}
}

// NewBot return a new Telegram Bot
func NewBot(cfg config.Config, ms messageService, us userService) Bot {
	b, _ := tgbotapi.NewBotAPI(cfg.Telegram.ApiToken)
	b.Debug = cfg.IsLog
	return &bot{
		api:            b,
		name:           cfg.Telegram.BotName,
		stat:           metrics.NewTelegramStat(),
		messageService: ms,
		userService:    us,
	}
}
