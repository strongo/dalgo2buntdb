package dalgo2buntdb

import (
	"context"
	"github.com/strongo/dalgo"
)

func (database) Select(ctx context.Context, query dalgo.Query) (dalgo.Reader, error) {
	return nil, nil
}
