package dalgo_buntdb

import (
	"context"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) RunInTransaction(ctx context.Context, f func(context.Context, dalgo.Transaction) error, options ...dalgo.TransactionOption) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return f(ctx, transaction{tx: tx})
	})
}

type transaction struct {
	tx *buntdb.Tx
}

func (t transaction) Upsert(_ context.Context, record dalgo.Record) error {
	panic("implement me")
}

var _ dalgo.Transaction = (*transaction)(nil)
