package dingtalk

import "github.com/tx7do/go-notifier"

type Option func(o *Notifier)

// Logger 日志记录器
func Logger(logger notifier.Logger) Option {
	return func(s *Notifier) {
		s.log = logger
	}
}

// AccessToken 钉钉访问令牌
func AccessToken(token string) Option {
	return func(s *Notifier) {
		s.accessToken = token
	}
}

// Secret 钉钉签名密钥
func Secret(secret string) Option {
	return func(s *Notifier) {
		s.secret = secret
	}
}
