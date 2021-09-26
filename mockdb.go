package dalgo2buntdb

import (
	"github.com/strongo/dalgo"
	"github.com/tidwall/buntdb"
	"testing"
)

// NewInMemoryMockDB creates a new dalgo.Database as a wrapper of of in-memory BuntDB instance
func NewInMemoryMockDB(t *testing.T) dalgo.Database {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatalf("Failed to create a new temporary in-memory mock BuntDB database: %v", err)
	}
	return NewDatabase(db)
}
