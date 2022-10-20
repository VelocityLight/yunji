package feishu

import (
	"testing"
	"yunji/configs"
)

func TestSendPostMsgCard(t *testing.T) {

	configs.LoadConfig("../../../config.yaml")
	config := configs.Config

	token, _ := GetAccessTokenFromApp(config.Feishu.AppId, config.Feishu.AppSecret)
	receiver := MsgReceiver{
		IDType: MsgIDTypeEmail,
		ID:     "yuchao.li@pingcap.com",
	}

	SendMsgCard(receiver,
		CardMsgWrapper{
			Msg: NewAlarmContentCard("😞 This is a alarm card test!", "**Alarm**: test [md](https://www.google.com)"),
		},
		token)
}
