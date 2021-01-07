package keeper

import (
	"github.com/faddat/testchain/x/testchain/types"
)

var _ types.QueryServer = Keeper{}
