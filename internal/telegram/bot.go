package telegram

import (
	"bot/internal/config"
	"bot/internal/system/metrics"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type bot struct {
	Api  *tgbotapi.BotAPI
	Name string
	Stat metrics.TelegramStat
}

// Bot is an interface which provides methods for bot working
type Bot interface {
	Listen()
}

// Listen messages
func (b *bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.Api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		fmt.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		b.Stat.IncReceivedMessages()
	}
}

// NewBot return a new Telegram Bot
func NewBot(cfg config.Config) Bot {
	b, _ := tgbotapi.NewBotAPI(cfg.Telegram.ApiToken)
	b.Debug = true
	return &bot{
		Api:  b,
		Name: cfg.Telegram.BotName,
		Stat: metrics.NewTelegramStat(),
	}
}
