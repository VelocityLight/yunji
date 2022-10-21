package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func MustConnect(dsn string) *Database {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(2048)
	db.SetMaxIdleConns(512)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(-1)

	return &Database{DB: db}
}

// Transaction will rollback if `f` returned any error.
func (s *Database) Transaction(ctx context.Context, f func(tx *sqlx.Tx) error, opts ...*sql.TxOptions) (err error) {
	var opt *sql.TxOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	tx, err := s.BeginTxx(ctx, opt)
	if err != nil {
		return err
	}
	needRollBack := true

	defer func() {
		if needRollBack {
			_ = tx.Rollback()
		}
	}()
	if err = f(tx); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	needRollBack = false
	return nil
}
