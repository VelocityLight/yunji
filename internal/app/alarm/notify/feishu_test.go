package notify

import (
	"testing"
	"yunji/configs"
	"yunji/internal/pkg/feishu"
)

func TestFeishuNotify(t *testing.T) {
	configs.LoadConfig("../../../../config.yaml")
	config := configs.Config
	feishu.SetFeishuApp(config.Feishu.AppId, config.Feishu.AppSecret)

	links := []Link{
		{
			"https://www.baidu.com",
			"百度",
		},
		{
			"https://www.google.com",
			"Google",
		},
	}

	block := Block{
		"这是一条**支持**<font color='red'>Markdown</font>的测试告警",
		links,
	}

	content := NotifyContent{
		"🛑 这是一条测试告警标题",
		[]Block{block, block},
	}

	notify := FeishuNotification{}
	notify.SendAlarm("yuchao.li@pingcap.com", content)
}
