package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Item struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Description string `json:"description" yaml:"description"`
    Price int32 `json:"price" yaml:"price"`
    Image string `json:"image" yaml:"image"`
}