package tsct

import (
	"os"
	"strconv"
)

type Config struct {
	BotToken   string
	WebhookURL string
	TelegramID int64
	QQSecret   string
}

var config *Config

func init() {
	Load(&Config{
		BotToken:   "1:test",
		WebhookURL: "",
		TelegramID: -1,
		QQSecret:   "bc106079e2e4789a5430261c4857de258459304502205164f9702ad6da2fd08",
	})
}

func Load(c *Config) {
	config = c
	LoadSecret(c.TelegramID, c.QQSecret)
	Bot(config.BotToken)

	if config.WebhookURL != "" {
		// 不推荐使用这里的功能来配置 Webhook
		Webhook(config.WebhookURL)
	}
}

func GetEnv() *Config {
	c := &Config{
		BotToken:   os.Getenv("BOT_TOKEN"),
		WebhookURL: os.Getenv("WEBHOOK_URL"),
		QQSecret:   os.Getenv("QQ_SECRET"),
	}
	c.TelegramID, _ = strconv.ParseInt(os.Getenv("TELEGRAM_ID"), 10, 64)
	if c.BotToken == "" {
		return nil
	}
	return c
}
