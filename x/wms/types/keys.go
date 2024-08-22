package types

const (
	// ModuleName defines the module name
	ModuleName = "wms"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wms"
)

var (
	ParamsKey = []byte("p_wms")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	WarehouseKey      = "Warehouse/value/"
	WarehouseCountKey = "Warehouse/count/"
)
