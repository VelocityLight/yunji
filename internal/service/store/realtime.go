package store

import (
	"context"
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

func (s *RealTimeService) GetServices() {}

func (s *RealTimeService) Create(ctx context.Context) error {
	return nil
}
