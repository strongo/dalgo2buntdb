package dalgo2buntdb

import (
	"github.com/strongo/dalgo/end2end"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	db := NewDatabase(openInMemoryDB(t))
	end2end.TestDalgoDB(t, db)
}
