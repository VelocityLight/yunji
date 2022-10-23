package notify

import (
	"fmt"
	"yunji/internal/pkg/feishu"
)

type FeishuNotification struct {
}

func (notify *FeishuNotification) SendAlarm(email string, content NotifyContent) error {
	token, err := feishu.GetAccessToken()
	if err != nil {
		return err
	}

	return feishu.SendMsgCard(
		feishu.MsgReceiver{
			IDType: feishu.MsgIDTypeEmail,
			ID:     email,
		},
		notify.Parse2Alarm(content),
		token,
	)
}

func (notify *FeishuNotification) Parse2Alarm(content NotifyContent) feishu.MsgWrapper {
	title := content.Header
	mdcontent := ""
	for _, block := range content.Blocks {
		mdcontent = mdcontent + fmt.Sprintf("%s \n", block.Text)
		for _, link := range block.Links {
			mdcontent = mdcontent + fmt.Sprintf("[%s](%s) \n", link.Text, link.Href)
		}

		mdcontent = mdcontent + "\n --------------\n"
	}

	card := feishu.NewAlarmContentCard(title, mdcontent)
	return feishu.CardMsgWrapper{
		Msg: card,
	}

}
