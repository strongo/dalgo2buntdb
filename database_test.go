package dalgo_buntdb

import (
	"github.com/tidwall/buntdb"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	var dtb = NewDatabase(db)
	if dtb == nil {
		t.Error("NewDatabase returned nil")
	}
}
