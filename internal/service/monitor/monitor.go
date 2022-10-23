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
			} else if v.Time > recentTime {
				recentTime = v.Time
			}
		}

		for k, v := range serviceCntPer10Sec {
			if v.avg == 0 {
				continue
			}
			if float64(summaryBy10Sec[recentTime][k]) > 5*v.avg {
				// Code for notification
				feishu := notify.FeishuNotification{}

				content := notify.NotifyContent{
					Header: fmt.Sprintf("🛑 Found hack in service **%s**", k),
				}

				// hard code email for demo
				feishu.SendAlarm("yejunchen@pingcap.com", content)
				feishu.SendAlarm("yuchao.li@pingcap.com", content)
			}
		}
	}

	return nil
}

func (m *Monitor) Shutdown() error {
	return nil
}
