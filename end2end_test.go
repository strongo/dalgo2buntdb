package dalgo_buntdb

import (
	e2e "github.com/strongo/dalgo-end2end-tests"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	db := NewDatabase(openInMemoryDB(t))
	e2e.EndToEnd(t, db)
}
