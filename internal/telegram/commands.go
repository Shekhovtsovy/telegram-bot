package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

const commandStart = "/start"
const commandAbout = "/about"

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
	b.log.Info("send message from bot", zap.String("command", commandStart))
	return nil
}

// handleCommandAbout handles about command
func (b *bot) handleCommandAbout(msg *tgbotapi.Message) error {
	aboutText := "*About*:\n\n Here your text..."
	answer := tgbotapi.NewMessage(msg.Chat.ID, aboutText)
	answer.ParseMode = "markdown"
	butRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🚀 Start", "|"+callbackStart),
	)
	answer.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(butRow)
	_, err := b.api.Send(answer)
	if err != nil {
		return err
	}
	b.log.Info("send message from bot", zap.String("command", commandAbout))
	return nil
}

// handleUnknownCommand handles unknown command
func (b *bot) handleUnknownCommand(msg *tgbotapi.Message) error {
	answer := tgbotapi.NewMessage(msg.Chat.ID, "Unknown command 🤷‍♀")
	_, err := b.api.Send(answer)
	if err != nil {
		return err
	}
	butRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🚀 Start", "|"+callbackStart),
	)
	butRow2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("❓ About", "|"+callbackAbout),
	)
	answer.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(butRow, butRow2)
	b.log.Info("send message from bot", zap.String("command", "unknown"))
	return nil
}
