package dalgo2buntdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

var ErrKeyAlreadyExists = errors.New("key already exists")

func (dtb database) Insert(ctx context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		t := transaction{tx: tx}
		return t.Insert(ctx, record, opts...)
	})
}

func (t transaction) Insert(ctx context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	options := dalgo.NewInsertOptions(opts...)
	generateID := options.IDGenerator()
	if generateID == nil {
		return t.insert(record)
	} else {
		return t.insertWithGenerator(ctx, generateID, record)
	}
}

func (t transaction) insertWithGenerator(ctx context.Context, generateID dalgo.IDGenerator, record dalgo.Record) error {
	for i := 0; i < 10; i++ {
		if err := generateID(ctx, record); err != nil {
			return err
		}
		if err := t.insert(record); err != nil {
			if err == ErrKeyAlreadyExists {
				continue
			}
			return err
		}
	}
	return nil
}

func (t transaction) insert(record dalgo.Record) error {
	key := record.Key()
	k := dalgo.GetRecordKeyPath(key)
	if _, err := t.tx.Get(k); err == nil {
		return ErrKeyAlreadyExists
	} else if err != buntdb.ErrNotFound {
		return err
	}
	s, err := json.Marshal(record.Data())
	if err != nil {
		return err
	}
	_, _, err = t.tx.Set(k, string(s), nil)
	return err
}
