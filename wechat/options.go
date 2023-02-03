package wechat

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

// Corp 访问令牌
func Corp(corpId, corpSecret string) Option {
	return func(s *Client) {
		s.corpId = corpId
		s.corpSecret = corpSecret
	}
}
