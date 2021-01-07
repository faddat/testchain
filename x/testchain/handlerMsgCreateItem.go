package testchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/faddat/testchain/x/testchain/types"
	"github.com/faddat/testchain/x/testchain/keeper"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Description: msg.Description,
    	Price: msg.Price,
    	Image: msg.Image,
	}
	k.CreateItem(ctx, item)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
