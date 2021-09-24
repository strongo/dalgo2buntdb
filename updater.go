package dalgo_buntdb

import (
	"context"
	"github.com/strongo/dalgo"
)

type updater struct {
	dtb *database
}

func (dtb database) Update(
	_ context.Context,
	_ *dalgo.Key,
	_ []dalgo.Update,
	_ ...dalgo.Precondition,
) error {
	panic("not supported")
}

func (dtb database) UpdateMulti(
	_ context.Context,
	_ []*dalgo.Key,
	_ []dalgo.Update,
	_ ...dalgo.Precondition,
) error {
	panic("not supported")
}
