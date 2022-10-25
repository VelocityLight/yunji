package monitor

import (
	"context"
	"fmt"
	"time"

	"yunji/common"
	"yunji/configs"
	"yunji/internal/app/alarm/notify"
	"yunji/internal/pkg/feishu"
	"yunji/internal/service/store"
)

type Monitor struct {
	Store  *store.Store
	Notify *notify.FeishuNotification
}

func NewMonitor(config *configs.ConfigYaml) *Monitor {
	feishu.SetFeishuApp(config.Feishu.AppId, config.Feishu.AppSecret)

	fmt.Print("init monitor\n")
	return &Monitor{
		Store:  store.NewStore(config),
		Notify: &notify.FeishuNotification{},
	}
}

func (m *Monitor) Inspect(_ context.Context) (err error) {
	var res common.GetRealTimeResponse
	for {
		time.Sleep(time.Second * 10)
		res, err = m.Store.RealTime.GetRealTime(context.Background())

		type serviceToCnt map[string]int
		var summaryBy10Sec = make(map[string]serviceToCnt)
		var serviceCntPer10Sec = make(map[string]struct {
			avg   float64
			total int
		})
		var recentTime string
		var subRecentTime string

		for _, v := range res.Body {
			serviceCntPer10Sec[v.Service] = struct {
				avg   float64
				total int
			}{
				avg:   (float64(serviceCntPer10Sec[v.Service].total)*serviceCntPer10Sec[v.Service].avg + float64(v.Cnt)) / float64(serviceCntPer10Sec[v.Service].total+1),
				total: serviceCntPer10Sec[v.Service].total + 1,
			}

			if _, ok := summaryBy10Sec[v.Time]; !ok {
				summaryBy10Sec[v.Time] = make(serviceToCnt)
			}
			summaryBy10Sec[v.Time][v.Service] += v.Cnt

			if recentTime == "" {
				recentTime = v.Time
				subRecentTime = recentTime
			} else if v.Time > recentTime {
				subRecentTime = recentTime
				recentTime = v.Time
			}
		}
		fmt.Print("here2\n")

		for k, v := range serviceCntPer10Sec {
			// fmt.Printf("%s, %v, %v\n", k, v, summaryBy10Sec[subRecentTime][k])

			if v.avg == 0 {
				continue
			}
			if float64(summaryBy10Sec[subRecentTime][k]) > 5*v.avg {
				// fmt.Printf("send alarm: found hack in service %s", k)

				// Code for notification
				feishu := notify.FeishuNotification{}

				block := notify.Block{
					Text: fmt.Sprintf("捕获到服务 **%s** 非正常创建，近十秒内创建出 <font color='red'>%d</font> 实例，严重超出正常范围，请关注", k, summaryBy10Sec[subRecentTime][k]),
				}

				content := notify.NotifyContent{
					Header: fmt.Sprintf("🛑 Found hack in service %s", k),
					Blocks: []notify.Block{block},
				}

				// hard code email for demo
				feishu.SendAlarm("xxx", content)
				feishu.SendAlarm("xxx", content)
			}
		}
	}

	return nil
}

func (m *Monitor) Shutdown() error {
	return nil
}
