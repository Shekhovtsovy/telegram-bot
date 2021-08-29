package config

type Telegram struct {
	ApiToken          string `envconfig:"TELEGRAM_API_TOKEN" default:""`
	BotName           string `envconfig:"TELEGRAM_BOT_NAME" default:""`
	IsLog             bool   `envconfig:"TELEGRAM_BOT_LOG" default:"false"`
	NeedToSubscribeOn string `envconfig:"TELEGRAM_NEED_TO_SUBSCRIBE_ON" default:""`
	BotId             int    `envconfig:"TELEGRAM_BOT_ID" default:""`
}
