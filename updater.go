package dalgo_buntdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) Update(
	ctx context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.Update(ctx, key, updates, preconditions...)
	})
}

func (dtb database) UpdateMulti(
	ctx context.Context,
	keys []*dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return transaction{tx: tx}.UpdateMulti(ctx, keys, updates, preconditions...)
	})
}

func (t transaction) Update(
	ctx context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	return t.update(ctx, key, updates, preconditions...)
}

func (t transaction) UpdateMulti(
	ctx context.Context,
	keys []*dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	for _, key := range keys {
		if err := t.update(ctx, key, updates, preconditions...); err != nil {
			return err
		}
	}
	return nil
}

func (t transaction) update(
	_ context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	k := dalgo.GetRecordKeyPath(key)
	s, err := t.tx.Get(k)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	if err = json.Unmarshal([]byte(s), &data); err != nil {
		return fmt.Errorf("failed to unmarshal data as JSON object: %v", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data as JSON object: %v", err)
	}
	for range updates {
	}
	_, _, err = t.tx.Set(k, string(b), nil)
	if err != nil {
		return err
	}
	return nil
}
