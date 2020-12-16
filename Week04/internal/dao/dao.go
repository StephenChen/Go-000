package dao

import (
	"context"
	"database/sql"
	"week04/internal/model"
)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	Coffee(c context.Context, id int64) (*model.Coffee, error)
}

type dao struct {
	db *sql.DB
}
