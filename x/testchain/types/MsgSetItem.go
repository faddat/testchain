package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetItem{}

type MsgSetItem struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Description string `json:"description" yaml:"description"`
  Price int32 `json:"price" yaml:"price"`
  Image string `json:"image" yaml:"image"`
}

func NewMsgSetItem(creator sdk.AccAddress, id string, description string, price int32, image string) MsgSetItem {
  return MsgSetItem{
    ID: id,
		Creator: creator,
    Description: description,
    Price: price,
    Image: image,
	}
}

func (msg MsgSetItem) Route() string {
  return RouterKey
}

func (msg MsgSetItem) Type() string {
  return "SetItem"
}

func (msg MsgSetItem) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetItem) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetItem) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}