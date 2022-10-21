package store

import "yunji/utils/sql"

func NewDBClient(dsn string) *sql.Database {
	return sql.MustConnect(dsn)
}

type Store struct {
	DB *sql.Database
}

func (s *Store) Shutdown() error {
	return s.DB.Close()
}
