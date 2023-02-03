package dingtalk

import "encoding/json"

type MsgType string

const (
	MsgTypeText       MsgType = "text"       // 文本消息
	MsgTypeMarkdown   MsgType = "markdown"   // markdown消息
	MsgTypeLink       MsgType = "link"       // 链接消息
	MsgTypeActionCard MsgType = "actionCard" // 卡片消息
	MsgTypeFeedCard   MsgType = "feedCard"   // 卡片消息
)

// Message interface
type Message interface {
	ToByte() ([]byte, error)
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Text struct {
	Content string `json:"content"`
}

type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"text"`
	At      At      `json:"at"`
}

func (m *TextMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeText
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// SetContent set content
func (m *TextMessage) SetContent(content string) *TextMessage {
	m.Text = Text{
		Content: content,
	}
	return m
}

// SetAt set at
func (m *TextMessage) SetAt(atMobiles []string, isAtAll bool) *TextMessage {
	m.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return m
}
