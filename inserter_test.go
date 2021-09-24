package dalgo_buntdb

import (
	"context"
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
	"testing"
)

const memory = ":memory:"

func TestInserter_Insert(t *testing.T) {
	bdb, err := buntdb.Open(memory)
	if err != nil {
		t.Fatalf("failed to open DB: %v", err)
	}
	ctx := context.Background()
	key := dalgo.NewKeyWithStrID("TestKind", "test-id")
	data := new(testKind)
	record := dalgo.NewRecord(key, data)
	db := NewDatabase(bdb)
	if err := db.Insert(ctx, record); err != nil {
		t.Errorf("expected to be successful, got error: %v", err)
	}
	if err := bdb.View(func(tx *buntdb.Tx) error {
		const id = "TestKind/test-id"
		if _, err := tx.Get(id); err != nil {
			t.Errorf("Inserted record is not found by ID: " + id)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
