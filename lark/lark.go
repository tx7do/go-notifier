package lark

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tx7do/go-notifier"
)

const larkApiUrl = "https://open.feishu.cn/open-apis/bot/v2/hook/"

type response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// Notifier 飞书客户端
type Notifier struct {
	log notifier.Logger

	cli *resty.Client

	accessToken string
	secret      string
}

func NewNotifier(opts ...Option) notifier.Notifier {
	c := &Notifier{
		log: notifier.DefaultLogger{},
	}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Notifier) init() {
	c.cli = resty.New()
}

// Send 发送聊天消息
func (c *Notifier) Send(_ context.Context, _, content string) error {
	if len(c.accessToken) < 1 {
		return errors.New("accessToken is empty")
	}

	timestamp := time.Now().Unix()
	sign, err := c.sign(c.secret, timestamp)
	if err != nil {
		return err
	}

	msg := &TextMessage{}
	msg.Content.Text = content

	body := msg.Body()
	body["timestamp"] = strconv.FormatInt(timestamp, 10)
	body["sign"] = sign

	_, err = json.Marshal(body)
	if err != nil {
		return err
	}

	URL := fmt.Sprintf("%v%v", larkApiUrl, c.accessToken)
	resp, err := c.cli.SetRetryCount(3).R().
		SetBody(body).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetResult(&response{}).
		ForceContentType("application/json").
		Post(URL)

	if err != nil {
		return err
	}

	result := resp.Result().(*response)
	if result.Code != 0 {
		return errors.New(result.Msg)
	}

	return nil
}

// sign 生成签名
func (c *Notifier) sign(secret string, timestamp int64) (string, error) {
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
