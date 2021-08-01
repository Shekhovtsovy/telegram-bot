package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

const commandStartRequest = "/start"
const commandStopRequest = "/stop"
const commandAboutRequest = "/about"
const commandAboutResponse = "📌 About bot \n\n Here your text..."

// Handle command Start
func (b *bot) handleCommandStart(msg *tgbotapi.Message) error {
	if b.cfg.NeedToSubscribeOn != "" {
		config := tgbotapi.ChatConfigWithUser{
			SuperGroupUsername: fmt.Sprintf("@%s", b.cfg.NeedToSubscribeOn),
			UserID:             msg.From.ID,
		}
		_, err := b.api.GetChatMember(config)
		if err != nil {
			b.log.Info("need to subscribe", zap.Int("userId", msg.From.ID))
			channelAddress := fmt.Sprintf("https://t.me/%s", b.cfg.NeedToSubscribeOn)
			answer := tgbotapi.NewMessage(msg.Chat.ID, "To continue you must subscribe: "+channelAddress)
			buttons := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(b.cfg.NeedToSubscribeOn, channelAddress))
			answer.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
			_, err := b.api.Send(answer)
			if err != nil {
				return err
			}
		}
	}
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
	_, err := b.api.Send(answer)
	if err != nil {
		return err
	}
	b.log.Info("send message from bot", zap.String("command", commandAboutRequest))
	return nil
}
