package dalgo2buntdb

import (
	"github.com/strongo/dalgo/dal"
	"github.com/tidwall/buntdb"
	"testing"
)

// NewInMemoryMockDB creates a new dal.Database as a wrapper of of in-memory BuntDB instance
func NewInMemoryMockDB(t *testing.T) dal.Database {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatalf("Failed to create a new temporary in-memory mock BuntDB database: %v", err)
	}
	return NewDatabase(db)
}
