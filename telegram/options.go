package telegram

import "github.com/tx7do/go-notifier"

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

// BotToken Telegram授权Token
func BotToken(token string) Option {
	return func(s *Notifier) {
		s.botToken = token
	}
}

// ChatId Telegram聊天ID
func ChatId(id int64) Option {
	return func(s *Notifier) {
		s.chatId = id
	}
}
