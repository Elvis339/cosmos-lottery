package keeper

import (
	"cosmos-lottery/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
