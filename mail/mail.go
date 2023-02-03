package mail

import (
	"github.com/go-gomail/gomail"
	"github.com/go-kratos/kratos/v2/log"
)

// Client 邮件客户端
type Client struct {
	log *log.Helper

	host     string
	port     int
	user     string
	password string

	mailer *gomail.Dialer

	toUsers []string
}

// NewClient 创建新的客户端
func NewClient(opts ...Option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Client) init() {
	c.mailer = gomail.NewDialer(c.host, c.port, c.user, c.password)
}

// Send 发送邮件
func (c *Client) Send(title, content string) error {
	sendMsg := gomail.NewMessage()
	sendMsg.SetHeader("From", c.user)
	sendMsg.SetHeader("To", c.toUsers...)
	sendMsg.SetHeader("Subject", title)
	sendMsg.SetBody("text/html", c.contentToHtml(content))
	err := c.mailer.DialAndSend(sendMsg)
	return err
}

// contentToHtml 讲发送内容转换为html格式
func (c *Client) contentToHtml(content string) string {
	return content
}
