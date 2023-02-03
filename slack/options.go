package slack

import (
	"github.com/tx7do/go-notifier"
)

type Option func(o *Notifier)

// Logger 日志记录器
func Logger(logger notifier.Logger) Option {
	return func(s *Notifier) {
		s.log = logger
	}
}

func Debug(debug bool) Option {
	return func(s *Notifier) {
		s.debug = debug
	}
}

// AuthToken Slack授权Token
func AuthToken(token string) Option {
	return func(s *Notifier) {
		s.authToken = token
	}
}

// ChannelId Slack频道ID
func ChannelId(id string) Option {
	return func(s *Notifier) {
		s.channelId = id
	}
}
