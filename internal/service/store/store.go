package store

import (
	"yunji/configs"
	"yunji/utils/sql"
)

func NewDBClient(dsn string) *sql.Database {
	return sql.MustConnect(dsn)
}

type Store struct {
	DB *sql.Database

	Billing  *BillingService
	RealTime *RealTimeService
}

func NewStore(config *configs.ConfigYaml) *Store {
	return &Store{
		DB: NewDBClient(config.Secret.DSN),

		Billing:  NewBillingService(NewDBClient(config.Secret.DSN)),
		RealTime: NewRealTimeService(NewDBClient(config.Secret.DSN)),
	}
}

func (s *Store) Shutdown() error {
	return s.DB.Close()
}
