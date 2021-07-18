package config

type Telegram struct {
	ApiToken string `envconfig:"TELEGRAM_API_TOKEN" default:""`
	BotName  string `envconfig:"TELEGRAM_BOT_NAME" default:""`
}
