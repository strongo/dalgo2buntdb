package dalgo2buntdb

import (
	"context"
	"github.com/strongo/dalgo/dal"
	"github.com/tidwall/buntdb"
)

func (dtb database) RunReadonlyTransaction(ctx context.Context, f dal.ROTxWorker, options ...dal.TransactionOption) error {
	return dtb.db.View(func(tx *buntdb.Tx) error {
		return f(ctx, transaction{tx: tx})
	})
}

func (dtb database) RunReadwriteTransaction(ctx context.Context, f dal.RWTxWorker, options ...dal.TransactionOption) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return f(ctx, transaction{tx: tx})
	})
}

type transaction struct {
	tx      *buntdb.Tx
	options dal.TransactionOptions
}

func (t transaction) Options() dal.TransactionOptions {
	return t.options
}

func (t transaction) Upsert(context.Context, dal.Record) error {
	panic("implement me")
}

func (t transaction) Select(context.Context, dal.Select) (dal.Reader, error) {
	panic("implement me")
}

var _ dal.Transaction = (*transaction)(nil)
