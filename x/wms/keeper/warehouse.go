package keeper

import (
	"context"
	"encoding/binary"

	"kewr/x/wms/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetWarehouseCount get the total number of warehouse
func (k Keeper) GetWarehouseCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.WarehouseCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWarehouseCount set the total number of warehouse
func (k Keeper) SetWarehouseCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.WarehouseCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWarehouse appends a warehouse in the store with a new id and update the count
func (k Keeper) AppendWarehouse(
	ctx context.Context,
	warehouse types.Warehouse,
) uint64 {
	// Create the warehouse
	count := k.GetWarehouseCount(ctx)

	// Set the ID of the appended value
	warehouse.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WarehouseKey))
	appendedValue := k.cdc.MustMarshal(&warehouse)
	store.Set(GetWarehouseIDBytes(warehouse.Id), appendedValue)

	// Update warehouse count
	k.SetWarehouseCount(ctx, count+1)

	return count
}

// SetWarehouse set a specific warehouse in the store
func (k Keeper) SetWarehouse(ctx context.Context, warehouse types.Warehouse) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WarehouseKey))
	b := k.cdc.MustMarshal(&warehouse)
	store.Set(GetWarehouseIDBytes(warehouse.Id), b)
}

// GetWarehouse returns a warehouse from its id
func (k Keeper) GetWarehouse(ctx context.Context, id uint64) (val types.Warehouse, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WarehouseKey))
	b := store.Get(GetWarehouseIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWarehouse removes a warehouse from the store
func (k Keeper) RemoveWarehouse(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WarehouseKey))
	store.Delete(GetWarehouseIDBytes(id))
}

// GetAllWarehouse returns all warehouse
func (k Keeper) GetAllWarehouse(ctx context.Context) (list []types.Warehouse) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WarehouseKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Warehouse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWarehouseIDBytes returns the byte representation of the ID
func GetWarehouseIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.WarehouseKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
