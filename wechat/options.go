package wechat

import "github.com/tx7do/go-notifier"

type Option func(o *Notifier)

// Logger 日志记录器
func Logger(logger notifier.Logger) Option {
	return func(s *Notifier) {
		s.log = logger
	}
}

// Corp 访问令牌
func Corp(corpId, corpSecret string) Option {
	return func(s *Notifier) {
		s.corpId = corpId
		s.corpSecret = corpSecret
	}
}
