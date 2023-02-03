package mail

import "github.com/tx7do/go-notifier"

type Option func(o *Notifier)

// Logger 日志记录器
func Logger(logger notifier.Logger) Option {
	return func(s *Notifier) {
		s.log = logger
	}
}

// Pop3 邮件服务器配置
func Pop3(host string, port int, user string, password string) Option {
	return func(c *Notifier) {
		c.host = host
		c.port = port
		c.user = user
		c.password = password
	}
}

// ToUsers 邮件接收人
func ToUsers(user ...string) Option {
	return func(c *Notifier) {
		c.toUsers = append(c.toUsers, user...)
	}
}
