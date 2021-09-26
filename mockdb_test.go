package dalgo2buntdb

import "testing"

func TestNewInMemoryMockDB(t *testing.T) {
	db := NewInMemoryMockDB(t)
	if db == nil {
		t.Fatal("db is null")
	}
}
