package testchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/faddat/testchain/x/testchain/types"
	"github.com/faddat/testchain/x/testchain/keeper"
)

func handleMsgSetItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Description: msg.Description,
    	Price: msg.Price,
    	Image: msg.Image,
	}
	if !msg.Creator.Equals(k.GetItemOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetItem(ctx, item)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
