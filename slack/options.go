package slack

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

// AuthToken Slack授权Token
func AuthToken(token string) Option {
	return func(s *Client) {
		s.authToken = token
	}
}

// ChannelId Slack频道ID
func ChannelId(id string) Option {
	return func(s *Client) {
		s.channelId = id
	}
}
