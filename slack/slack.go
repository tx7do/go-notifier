package slack

import (
	"context"
	"time"

	"github.com/slack-go/slack"
	"github.com/tx7do/go-notifier"
)

// Notifier Slack客户端
type Notifier struct {
	log notifier.Logger

	cli *slack.Client

	debug     bool
	authToken string
	channelId string
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
	c.cli = slack.New(c.authToken, slack.OptionDebug(c.debug))
}

// Send 发送聊天消息
func (c *Notifier) Send(_ context.Context, title, content string) error {
	attachment := slack.Attachment{
		Pretext: title,
		Text:    content,
		Color:   "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := c.cli.PostMessage(
		c.channelId,
		slack.MsgOptionAttachments(attachment),
	)

	c.log.Info("Message sent at %s", timestamp)

	return err
}
