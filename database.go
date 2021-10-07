package dalgo2buntdb

import (
	"context"
	"github.com/strongo/dalgo/dal"
	"github.com/tidwall/buntdb"
)

type database struct {
	db *buntdb.DB
}

var _ dal.Database = (*database)(nil)

// NewDatabase creates a new instance of DALgo adapter for BungDB
func NewDatabase(db *buntdb.DB) dal.Database {
	if db == nil {
		panic("db is a required parameter, got nil")
	}
	return database{
		db: db,
	}
}

func (dtb database) Upsert(ctx context.Context, record dal.Record) error {
	panic("implement me")
}
