package dalgo_buntdb

import (
	"context"
	"encoding/json"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) Get(ctx context.Context, record dalgo.Record) error {
	return dtb.db.View(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.Get(ctx, record)
	})
}

func (dtb database) GetMulti(ctx context.Context, records []dalgo.Record) error {
	return dtb.db.View(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.GetMulti(ctx, records)
	})
}

func (t transaction) Get(_ context.Context, record dalgo.Record) error {
	key := record.Key()
	keyPath := dalgo.GetRecordKeyPath(key)
	s, err := t.tx.Get(keyPath)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(s), record.Data())
}

func (t transaction) GetMulti(ctx context.Context, records []dalgo.Record) error {
	for _, rec := range records {
		keyPath := dalgo.GetRecordKeyPath(rec.Key())
		s, err := t.tx.Get(keyPath)
		if err != nil {
			return err
		}
		return json.Unmarshal([]byte(s), rec.Data())
	}
	return nil
}
