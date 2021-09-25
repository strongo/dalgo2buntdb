package dalgo2buntdb

import (
	"context"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) Delete(ctx context.Context, key *dalgo.Key) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.Delete(ctx, key)
	})
}

func (dtb database) DeleteMulti(_ context.Context, keys []*dalgo.Key) (err error) {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		for _, key := range keys {
			keyPath := dalgo.GetRecordKeyPath(key)
			if _, err = tx.Delete(keyPath); err != nil {
				return err
			}
		}
		return err
	})
}

func (t transaction) Delete(ctx context.Context, key *dalgo.Key) error {
	keyPath := dalgo.GetRecordKeyPath(key)
	_, err := t.tx.Delete(keyPath)
	return err
}

func (t transaction) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	for _, key := range keys {
		if err := t.Delete(ctx, key); err != nil {
			return err
		}
	}
	return nil
}
