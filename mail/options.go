package mail

import "github.com/go-kratos/kratos/v2/log"

type Option func(o *Client)

// Logger 日志记录器
func Logger(logger log.Logger) Option {
	return func(c *Client) {
		c.log = log.NewHelper(logger)
	}
}

// Pop3 邮件服务器配置
func Pop3(host string, port int, user string, password string) Option {
	return func(c *Client) {
		c.host = host
		c.port = port
		c.user = user
		c.password = password
	}
}

// ToUsers 邮件接收人
func ToUsers(user ...string) Option {
	return func(c *Client) {
		c.toUsers = append(c.toUsers, user...)
	}
}
