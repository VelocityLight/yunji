package store

import (
	"context"
	"fmt"

	"yunji/common"
	"yunji/utils/sql"
)

type RealTimeService struct {
	db *sql.Database
}

func NewRealTimeService(db *sql.Database) *RealTimeService {
	return &RealTimeService{db}
}

func (s *RealTimeService) Update() {}

func (s *RealTimeService) Get() {}

func (s *RealTimeService) Delete() {}

func (s *RealTimeService) List(ctx context.Context, opts common.QueryRealtimeEventOpts) ([]common.RealtimeEvent, error) {
	var events []common.RealtimeEvent
	return events, nil
}

func (s *RealTimeService) GetRealTime(ctx context.Context) (common.GetRealTimeResponse, error) {
	var res common.GetRealTimeResponse
	err := s.db.SelectContext(ctx, &res.Body, `
		SELECT 
			DATE_FORMAT(
				concat( 
					date( created_time ), ' ', 
					HOUR ( created_time ), ':', 
					MINUTE(created_time), ':', 
					floor( SECOND ( created_time ) / 10 ) * 10 
				),
				'%Y-%m-%d %H:%i:%S'
			) AS time,
			count(created_time) AS cnt,
			product_code AS service 
		FROM realtime_event WHERE created_time > DATE_ADD(NOW(),INTERVAL 470 MINUTE)
		GROUP BY time,service
		ORDER BY time`)
	return res, err
}

func (s *RealTimeService) GetRealTimeForMonitor(ctx context.Context) (common.GetRealTimeResponse, error) {
	var res common.GetRealTimeResponse
	err := s.db.SelectContext(ctx, &res.Body, `
		SELECT 
			DATE_FORMAT(
				concat( 
					date( created_time ), ' ', 
					HOUR ( created_time ), ':', 
					MINUTE(created_time), ':', 
					floor( SECOND ( created_time ) / 10 ) * 10 
				),
				'%Y-%m-%d %H:%i:%S'
			) AS time,
			count(created_time) AS cnt,
			product_code AS service 
		FROM realtime_event WHERE created_time > DATE_ADD(NOW(),INTERVAL 475 MINUTE)
		GROUP BY time,service
		ORDER BY time`)
	return res, err
}

func (s *RealTimeService) Create(event common.RealtimeEvent) error {
	insertString := fmt.Sprintf(`
    INSERT INTO realtime_event (event_id, account_id, product_code, product_name,
        product_region, resource_id, created_time, usage_type, operation, used_by)
    VALUES ('%s', '%s', '%s', '%s', '%s','%s','%s','%s','%s','%s');
    `, event.EventID, event.AccountID, event.ProductCode, event.ProductName,
		event.ProductRegion, event.ResourceID, event.CreatedTime.Format("2006-01-02 15:04:05"),
		event.UsageType, event.Operation,
		event.UsedByTag)
	_, err := s.db.Exec(insertString)
	return err
}
