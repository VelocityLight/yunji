package notify

import (
	"fmt"
	"testing"
	"yunji/configs"
	"yunji/internal/pkg/feishu"
)

func TestFeishuNotify(t *testing.T) {
	configs.LoadConfig("../../../../config.yaml")
	config := configs.Config
	feishu.SetFeishuApp(config.Feishu.AppId, config.Feishu.AppSecret)

	block := Block{
		Text: "这是一条**支持**<font color='red'>Markdown</font>的测试告警",
	}

	content := NotifyContent{
		"🛑 这是一条测试告警标题",
		[]Block{block, block},
	}

	notify := FeishuNotification{}
	notify.SendAlarm("xxxxxxxxxxxxxxxxxxxxx", content)

}

func TestFeishuNotify1(t *testing.T) {
	configs.LoadConfig("../../../../config.yaml")
	config := configs.Config
	feishu.SetFeishuApp(config.Feishu.AppId, config.Feishu.AppSecret)

	// Code for notification
	feishu := FeishuNotification{}

	content := NotifyContent{
		Header: fmt.Sprintf("🛑 Found hack in service **%s**", "xxx"),
	}

	// hard code email for demo
	feishu.SendAlarm("xxx", content)
	feishu.SendAlarm("xxx", content)

}
