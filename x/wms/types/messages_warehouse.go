package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateWarehouse{}

func NewMsgCreateWarehouse(creator string, name string, location string, capacity uint64) *MsgCreateWarehouse {
	return &MsgCreateWarehouse{
		Creator:  creator,
		Name:     name,
		Location: location,
		Capacity: capacity,
	}
}

func (msg *MsgCreateWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWarehouse{}

func NewMsgUpdateWarehouse(creator string, id uint64, name string, location string, capacity uint64) *MsgUpdateWarehouse {
	return &MsgUpdateWarehouse{
		Id:       id,
		Creator:  creator,
		Name:     name,
		Location: location,
		Capacity: capacity,
	}
}

func (msg *MsgUpdateWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWarehouse{}

func NewMsgDeleteWarehouse(creator string, id uint64) *MsgDeleteWarehouse {
	return &MsgDeleteWarehouse{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
