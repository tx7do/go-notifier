package slack

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/slack-go/slack"
)

// Client Slack客户端
type Client struct {
	log *log.Helper

	cli *slack.Client

	debug     bool
	authToken string
	channelId string
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
	c.cli = slack.New(c.authToken, slack.OptionDebug(c.debug))
}

// Send 发送聊天消息
func (c *Client) Send(title, content string) error {
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

	c.log.Infof("Message sent at %s", timestamp)

	return err
}
