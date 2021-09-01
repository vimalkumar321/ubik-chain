package keeper

import (
	"github.com/example/ubik/x/ubik/types"
)

var _ types.QueryServer = Keeper{}
