package dalgo_buntdb

import (
	"context"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) Delete(_ context.Context, key *dalgo.Key) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		keyPath := dalgo.GetRecordKeyPath(key)
		_, err := tx.Delete(keyPath)
		return err
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
