package feishu

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Option func(o *Client)

// Logger 日志记录器
func Logger(logger log.Logger) Option {
	return func(s *Client) {
		s.log = log.NewHelper(logger)
	}
}

// AccessToken 飞书访问令牌
func AccessToken(token string) Option {
	return func(s *Client) {
		s.accessToken = token
	}
}

// Secret 飞书签名密钥
func Secret(secret string) Option {
	return func(s *Client) {
		s.secret = secret
	}
}
