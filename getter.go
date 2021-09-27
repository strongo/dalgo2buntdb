package dalgo2buntdb

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
	keyPath := key.String()
	s, err := t.tx.Get(keyPath)
	if err != nil {
		if err == buntdb.ErrNotFound {
			err = dalgo.NewErrNotFoundByKey(key, err)
		}
		return err
	}
	return json.Unmarshal([]byte(s), record.Data())
}

func (t transaction) GetMulti(ctx context.Context, records []dalgo.Record) error {
	for _, record := range records {
		key := record.Key()
		keyPath := key.String()
		s, err := t.tx.Get(keyPath)
		if err == buntdb.ErrNotFound {
			record.SetError(dalgo.NewErrNotFoundByKey(key, err))
			continue
		} else if err != nil {
			return err
		}
		record.SetError(err)
		if err = json.Unmarshal([]byte(s), record.Data()); err != nil {
			record.SetError(err)
		}
	}
	return nil
}
