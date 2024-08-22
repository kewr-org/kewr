package keeper

import (
	"context"
	"fmt"

	"kewr/x/wms/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateWarehouse(goCtx context.Context, msg *types.MsgCreateWarehouse) (*types.MsgCreateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var warehouse = types.Warehouse{
		Creator:  msg.Creator,
		Name:     msg.Name,
		Location: msg.Location,
		Capacity: msg.Capacity,
	}

	id := k.AppendWarehouse(
		ctx,
		warehouse,
	)

	return &types.MsgCreateWarehouseResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateWarehouse(goCtx context.Context, msg *types.MsgUpdateWarehouse) (*types.MsgUpdateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var warehouse = types.Warehouse{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Name:     msg.Name,
		Location: msg.Location,
		Capacity: msg.Capacity,
	}

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWarehouse(ctx, warehouse)

	return &types.MsgUpdateWarehouseResponse{}, nil
}

func (k msgServer) DeleteWarehouse(goCtx context.Context, msg *types.MsgDeleteWarehouse) (*types.MsgDeleteWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWarehouse(ctx, msg.Id)

	return &types.MsgDeleteWarehouseResponse{}, nil
}
