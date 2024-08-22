package keeper

import (
	"context"

	"kewr/x/wms/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WarehouseAll(ctx context.Context, req *types.QueryAllWarehouseRequest) (*types.QueryAllWarehouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var warehouses []types.Warehouse

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	warehouseStore := prefix.NewStore(store, types.KeyPrefix(types.WarehouseKey))

	pageRes, err := query.Paginate(warehouseStore, req.Pagination, func(key []byte, value []byte) error {
		var warehouse types.Warehouse
		if err := k.cdc.Unmarshal(value, &warehouse); err != nil {
			return err
		}

		warehouses = append(warehouses, warehouse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWarehouseResponse{Warehouse: warehouses, Pagination: pageRes}, nil
}

func (k Keeper) Warehouse(ctx context.Context, req *types.QueryGetWarehouseRequest) (*types.QueryGetWarehouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	warehouse, found := k.GetWarehouse(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetWarehouseResponse{Warehouse: warehouse}, nil
}
