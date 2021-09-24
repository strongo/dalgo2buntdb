package dalgo_buntdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
)

func (dtb database) RunInTransaction(ctx context.Context, f func(context.Context, dalgo.Transaction) error, options ...dalgo.TransactionOption) error {
	firestoreTxOptions := createFirestoreTransactionOptions(options)
	return dtb.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return f(ctx, transaction{tx: tx})
	}, firestoreTxOptions...)
}

func createFirestoreTransactionOptions(opts []dalgo.TransactionOption) (options []firestore.TransactionOption) {
	to := dalgo.NewTransactionOptions(opts...)
	if to.IsReadonly() {
		options = append(options, firestore.ReadOnly)
	}
	return
}

type transaction struct {
	tx  *buntdb.Tx
}


func (t transaction) Upsert(_ context.Context, record dalgo.Record) error {
	dr := t.dtb.doc(record.Key())
	return t.tx.Set(dr, record.Data())
}




func (t transaction) Update(
	_ context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	dr := t.dtb.doc(key)
	fsUpdates := make([]firestore.Update, len(updates))
	for i, u := range updates {
		fsUpdates[i] = getFirestoreUpdate(u)
	}
	fsPreconditions := getUpdatePreconditions(preconditions)
	return t.tx.Update(dr, fsUpdates, fsPreconditions...)
}

func getFirestoreUpdate(update dalgo.Update) firestore.Update {
	return firestore.Update{
		Path:      update.Field,
		FieldPath: (firestore.FieldPath)(update.FieldPath),
		Value:     update.Value,
	}
}

func getUpdatePreconditions(preconditions []dalgo.Precondition) (fsPreconditions []firestore.Precondition) {
	p := dalgo.GetPreconditions(preconditions...)
	if p.Exists() {
		fsPreconditions = []firestore.Precondition{firestore.Exists}
	}
	return fsPreconditions
}

func (t transaction) Delete(ctx context.Context, key *dalgo.Key) error {
	dr := t.dtb.doc(key)
	return t.tx.Delete(dr)
}




func (t transaction) UpdateMulti(
	_ context.Context,
	keys []*dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	fsPreconditions := getUpdatePreconditions(preconditions)
	for _, key := range keys {
		dr := t.dtb.doc(key)
		fsUpdates := make([]firestore.Update, len(updates))
		for i, u := range updates {
			fsUpdates[i] = getFirestoreUpdate(u)
		}
		if err := t.tx.Update(dr, fsUpdates, fsPreconditions...); err != nil {
			keyPath := dalgo.GetRecordKeyPath(key)
			return fmt.Errorf("failed to update record with key: %v: %w", keyPath, err)
		}
	}
	return nil
}

func (t transaction) DeleteMulti(_ context.Context, keys []*dalgo.Key) error {
	for _, k := range keys {
		dr := t.dtb.doc(k)
		if err := t.tx.Delete(dr); err != nil {
			return fmt.Errorf("failed to delete record: %w", err)
		}
	}
	return nil
}

var _ dalgo.Transaction = (*transaction)(nil)
