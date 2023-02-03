package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tx7do/go-notifier"
)

// Notifier Telegram客户端
type Notifier struct {
	log notifier.Logger

	cli *tgbotapi.BotAPI

	debug    bool
	botToken string
	chatId   int64
}

func NewNotifier(opts ...Option) notifier.Notifier {
	c := &Notifier{
		debug: false,
		log:   notifier.DefaultLogger{},
	}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Notifier) init() {
	c.cli, _ = tgbotapi.NewBotAPI(c.botToken)
	c.cli.Debug = c.debug
}

// Send 发送聊天消息
func (c *Notifier) Send(_ context.Context, _, content string) error {

	msg := tgbotapi.NewMessage(c.chatId, content)

	if _, err := c.cli.Send(msg); err != nil {
		return err
	}

	return nil
}
