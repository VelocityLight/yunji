package feishu

import "encoding/json"

type MsgWrapper interface {
	GetMsgType() string
	GetMsgJson() string
}

type TextMsgWrapper struct {
	MsgType string `json:"msg_type"`
	Msg     string `json:"content"`
}

func (msg TextMsgWrapper) GetMsgType() string {
	return "text"
}

func (msg TextMsgWrapper) GetMsgJson() string {
	return msg.Msg
}

var _ MsgWrapper = TextMsgWrapper{}

type PostMsgWrapper struct {
	MsgType string      `json:"msg_type"`
	Msg     ContentPost `json:"content"`
}

func (msg PostMsgWrapper) GetMsgType() string {
	return "post"
}

func (msg PostMsgWrapper) GetMsgJson() string {
	jsonData, _ := json.Marshal(msg.Msg)
	return string(jsonData)
}

var _ MsgWrapper = PostMsgWrapper{}

type ContentPost struct {
	ZhCnContent ContentWrapper `json:"zh_cn"`
	// Post PostCN `json:"post,omitempty"`
}

type PostCN struct {
	ZhCnContent ContentWrapper `json:"zh_cn"`
}

type ContentWrapper struct {
	Title string             `json:"title"`
	Rows  [][]ContentElement `json:"content"`
}

type ContentElement interface {
}

type TextContentElement struct {
	Tag      string `json:"tag" default:"text"`
	UnEscape bool   `json:"un_escape,omitempty" default:"false"`
	Text     string `json:"text"`
}

func NewTextContentElement(text string) TextContentElement {
	return TextContentElement{
		Tag:  "text",
		Text: text,
	}
}

var _ ContentElement = TextContentElement{}

type HrefContentElement struct {
	Tag      string `json:"tag" default:"a"`
	UnEscape bool   `json:"un_escape,omitempty" default:"false"`
	Href     string `json:"href"`
	Text     string `json:"text"`
}

func NewHrefContentElement(href, text string) HrefContentElement {
	return HrefContentElement{
		Tag:  "a",
		Href: href,
		Text: text,
	}
}

var _ ContentElement = HrefContentElement{}

type CardMsgWrapper struct {
	MsgType string      `json:"msg_type"`
	Msg     ContentCard `json:"content"`
}

func (msg CardMsgWrapper) GetMsgType() string {
	return "interactive"
}

func (msg CardMsgWrapper) GetMsgJson() string {
	jsonData, _ := json.Marshal(msg.Msg)
	return string(jsonData)
}

var _ MsgWrapper = PostMsgWrapper{}

type ContentCard struct {
	Config   CardConfig    `json:"config"`
	Header   CardHeader    `json:"header"`
	Elements []CardElement `json:"elements"`
}

type CardConfig struct {
	UpdateMulti    bool `json:"update_multi"`
	WideScreenMode bool `json:"wide_screen_mode"`
}

func NewCardConfig(updateMulti, wideScreenMode bool) CardConfig {
	return CardConfig{
		UpdateMulti:    updateMulti,
		WideScreenMode: wideScreenMode,
	}
}

type CardHeader struct {
	Title    CardHeaderTitle `json:"title"`
	Template string          `json:"template"`
}

type CardHeaderTitle struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func NewCardHeader(title, template string) CardHeader {
	headerTitle := CardHeaderTitle{
		Tag:     "plain_text",
		Content: title,
	}

	return CardHeader{
		Title:    headerTitle,
		Template: template,
	}
}

type CardElement interface{}

type cardElementMD struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func NewMDCardElement(mdcontent string) cardElementMD {
	return cardElementMD{
		Tag:     "markdown",
		Content: mdcontent,
	}
}

func NewAlarmContentCard(title, mdcontent string) ContentCard {
	config := NewCardConfig(true, true)
	header := NewCardHeader(title, "red")
	content := NewMDCardElement(mdcontent)
	element := []CardElement{content}

	return ContentCard{
		Config:   config,
		Header:   header,
		Elements: element,
	}
}
