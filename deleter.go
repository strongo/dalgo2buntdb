package dalgo2buntdb

import (
	"context"
	"github.com/strongo/dalgo/dal"
	"github.com/tidwall/buntdb"
)

func (dtb database) Delete(ctx context.Context, key *dal.Key) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		err := transaction{tx: tx}.Delete(ctx, key)
		if err == buntdb.ErrNotFound {
			err = nil
		}
		return err
	})
}

func (dtb database) DeleteMulti(_ context.Context, keys []*dal.Key) (err error) {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		for _, key := range keys {
			keyPath := key.String()
			if _, err = tx.Delete(keyPath); err != nil {
				if err == buntdb.ErrNotFound {
					err = nil
					continue
				}
				return err
			}
		}
		return err
	})
}

func (t transaction) Delete(ctx context.Context, key *dal.Key) error {
	keyPath := key.String()
	_, err := t.tx.Delete(keyPath)
	return err
}

func (t transaction) DeleteMulti(ctx context.Context, keys []*dal.Key) error {
	for _, key := range keys {
		if err := t.Delete(ctx, key); err != nil {
			return err
		}
	}
	return nil
}
