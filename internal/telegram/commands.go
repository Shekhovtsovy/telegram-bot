package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

const commandStartRequest = "/start"
const commandStopRequest = "/stop"
const commandAboutRequest = "/about"
const commandAboutResponse = "📌 About bot \n\n Here your text..."

// Handle command Start
func (b *bot) handleCommandStart(msg *tgbotapi.Message) error {
	b.log.Info("send message from bot", zap.String("command", commandStartRequest))
	return nil
}

// Handle command Stop
func (b *bot) handleCommandStop(msg *tgbotapi.Message) error {
	b.log.Info("send message from bot", zap.String("command", commandStopRequest))
	return nil
}

// Handle command About
func (b *bot) handleCommandAbout(msg *tgbotapi.Message) error {
	answer := tgbotapi.NewMessage(msg.Chat.ID, commandAboutResponse)
	buttons := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Button 1", "Callback 1"), tgbotapi.NewInlineKeyboardButtonData("Button 2", "Callback 2"))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)
	answer.ReplyMarkup = keyboard
	_, err := b.api.Send(answer)
	if err != nil {
		return err
	}
	b.log.Info("send message from bot", zap.String("command", commandAboutRequest))
	return nil
}
