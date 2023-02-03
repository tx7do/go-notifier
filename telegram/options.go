package telegram

import "github.com/go-kratos/kratos/v2/log"

type Option func(o *Client)

// Logger 日志记录器
func Logger(logger log.Logger) Option {
	return func(s *Client) {
		s.log = log.NewHelper(logger)
	}
}

func Debug(debug bool) Option {
	return func(s *Client) {
		s.debug = debug
	}
}

// BotToken Telegram授权Token
func BotToken(token string) Option {
	return func(s *Client) {
		s.botToken = token
	}
}

// ChatId Telegram聊天ID
func ChatId(id int64) Option {
	return func(s *Client) {
		s.chatId = id
	}
}
