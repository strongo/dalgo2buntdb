package dalgo_buntdb

import (
	"context"
	"encoding/json"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) Set(ctx context.Context, record dalgo.Record) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.Set(ctx, record)
	})
}

func (dtb database) SetMulti(ctx context.Context, records []dalgo.Record) error {
	batch := dtb.batch()
	for _, record := range records {
		key := record.Key()
		docRef := dtb.doc(key)
		data := record.Data()
		batch.Set(docRef, data)
	}
	_, err := batch.Commit(ctx)
	return err
}

func (t transaction) Set(ctx context.Context, record dalgo.Record) error {
	key := record.Key()
	k := dalgo.GetRecordKeyPath(key)
	s, err := json.Marshal(record.Data())
	if err != nil {
		return err
	}
	_, _, err = t.tx.Set(k, s, nil)
	return err
}

func (t transaction) SetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}
