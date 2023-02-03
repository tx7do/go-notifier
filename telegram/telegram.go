package telegram

import (
	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client Telegram客户端
type Client struct {
	log *log.Helper

	cli *tgbotapi.BotAPI

	debug    bool
	botToken string
	chatId   int64
}

func NewClient(opts ...Option) *Client {
	c := &Client{
		debug: false,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Client) init() {
	c.cli, _ = tgbotapi.NewBotAPI(c.botToken)
	c.cli.Debug = c.debug
}

// Send 发送聊天消息
func (c *Client) Send(_, content string) error {

	msg := tgbotapi.NewMessage(c.chatId, content)

	if _, err := c.cli.Send(msg); err != nil {
		return err
	}

	return nil
}
