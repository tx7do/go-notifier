package wechat

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
)

const (
	sendMessageUrl = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	getTokenUrl    = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

type response struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (a *accessToken) Valid() bool {
	return a.AccessToken != ""
}

func (a *accessToken) Expires() bool {
	return a.ExpiresIn > 0
}

type textMessage struct {
	ToUser  string            `json:"touser"`
	ToParty string            `json:"toparty"`
	ToTag   string            `json:"totag"`
	MsgType string            `json:"msgtype"`
	AgentId int               `json:"agentid"`
	Text    map[string]string `json:"text"`
	Safe    int               `json:"safe"`
}

// Client 微信客户端
// https://cloud.tencent.com/developer/news/194102
// https://www.daimajiaoliu.com/daima/487120fbb9003f4
// https://www.cxyzjd.com/article/suiban7403/78198536
type Client struct {
	log *log.Helper

	cli *resty.Client

	corpId     string
	corpSecret string

	token accessToken
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

	if !c.token.Valid() || c.token.Expires() {
		token, err := c.getToken(c.corpId, c.corpSecret)
		if err != nil {
			return err
		}
		c.token = token
	}

	//body := bytes.NewBuffer([]byte(content))

	URL := fmt.Sprintf("%v%v", sendMessageUrl, c.token.AccessToken)
	resp, err := c.cli.SetRetryCount(3).R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		ForceContentType("application/json").
		SetBody(content).
		SetResult(&response{}).
		Post(URL)

	if err != nil {
		return err
	}

	result := resp.Result().(*response)
	if result.Code != 0 && result.Msg != "ok" {
		return errors.New(result.Msg)
	}

	return nil
}

func (c *Client) getToken(corpId, corpSecret string) (at accessToken, err error) {
	URL := getTokenUrl + corpId + "&corpsecret=" + corpSecret
	resp, err := c.cli.SetRetryCount(3).R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		ForceContentType("application/json").
		SetResult(&accessToken{}).
		Get(URL)

	if err != nil {
		return at, err
	}

	at = *resp.Result().(*accessToken)
	if !at.Valid() {
		err = errors.New("corpid or corpsecret error")
	}
	return
}
