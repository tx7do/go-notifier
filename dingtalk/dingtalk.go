package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"math"
	"net/url"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// https://oapi.dingtalk.com/robot/send?access_token=xxx
const dingTalkOAPI = "oapi.dingtalk.com"

var dingTalkURL url.URL = url.URL{
	Scheme: "https",
	Host:   dingTalkOAPI,
	Path:   "robot/send",
}

type response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

// Client 钉钉客户端
type Client struct {
	log *log.Helper

	cli *resty.Client

	accessToken string
	secret      string
}

func NewClient(opts ...Option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	c.init()
	return c
}

// init 初始化
func (c *Client) init() {
	c.cli = resty.New()
}

// Send 发送聊天消息
func (c *Client) Send(_, content string) error {

	pushURL, err := c.url(c.accessToken, c.secret)
	if err != nil {
		return err
	}

	msg := &TextMessage{}
	msg.SetContent(content)

	reqBytes, err := msg.ToByte()
	if err != nil {
		return err
	}

	resp, err := c.cli.SetRetryCount(3).R().
		SetBody(reqBytes).
		SetHeader("Accept", "application/json").
		SetHeader("Accept-Charset", "utf8").
		SetHeader("Content-Type", "application/json").
		SetResult(&response{}).
		ForceContentType("application/json").
		Post(pushURL)

	if err != nil {
		return err
	}

	result := resp.Result().(*response)

	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}

	return err
}

// 获取钉钉API调用链接
// 如果没有加签，secret 设置为 "" 即可
func (c *Client) url(accessToken string, secret string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	return c.urlWithTimestamp(timestamp, accessToken, secret)
}

// urlWithTimestamp 获取钉钉API调用链接
func (c *Client) urlWithTimestamp(timestamp string, accessToken string, secret string) (string, error) {
	dtu := dingTalkURL
	value := url.Values{}
	value.Set("access_token", accessToken)

	if secret == "" {
		dtu.RawQuery = value.Encode()
		return dtu.String(), nil
	}

	sign, err := c.sign(timestamp, secret)
	if err != nil {
		dtu.RawQuery = value.Encode()
		return dtu.String(), err
	}

	value.Set("timestamp", timestamp)
	value.Set("sign", sign)
	dtu.RawQuery = value.Encode()
	return dtu.String(), nil
}

// sign 生成签名
func (c *Client) sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// validate 校验请求是否合法
// https://ding-doc.dingtalk.com/doc#/serverapi2/elzz1p
func (c *Client) validate(signStr, timestamp, secret string) (bool, error) {
	t, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false, err
	}

	timeGap := time.Since(time.Unix(t, 0))
	if math.Abs(timeGap.Hours()) > 1 {
		return false, fmt.Errorf("specified timestamp is expired")
	}

	ourSign, err := c.sign(timestamp, secret)
	if err != nil {
		return false, err
	}
	return ourSign == signStr, nil
}
