package mail

import (
	"context"

	"github.com/go-gomail/gomail"
	"github.com/tx7do/go-notifier"
)

// Notifier 邮件客户端
type Notifier struct {
	log notifier.Logger

	host     string
	port     int
	user     string
	password string

	mailer *gomail.Dialer

	toUsers []string
}

// NewNotifier 创建新的客户端
func NewNotifier(opts ...Option) notifier.Notifier {
	c := &Notifier{
		log: notifier.DefaultLogger{},
	}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Notifier) init() {
	c.mailer = gomail.NewDialer(c.host, c.port, c.user, c.password)
}

// Send 发送邮件
func (c *Notifier) Send(_ context.Context, title, content string) error {
	sendMsg := gomail.NewMessage()
	sendMsg.SetHeader("From", c.user)
	sendMsg.SetHeader("To", c.toUsers...)
	sendMsg.SetHeader("Subject", title)
	sendMsg.SetBody("text/html", c.contentToHtml(content))
	err := c.mailer.DialAndSend(sendMsg)
	return err
}

// contentToHtml 讲发送内容转换为html格式
func (c *Notifier) contentToHtml(content string) string {
	return content
}
