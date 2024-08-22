package wms

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kewr/api/kewr/wms"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "WarehouseAll",
					Use:       "list-warehouse",
					Short:     "List all Warehouse",
				},
				{
					RpcMethod:      "Warehouse",
					Use:            "show-warehouse [id]",
					Short:          "Shows a Warehouse by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateWarehouse",
					Use:            "create-warehouse [name] [location] [capacity]",
					Short:          "Create Warehouse",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "location"}, {ProtoField: "capacity"}},
				},
				{
					RpcMethod:      "UpdateWarehouse",
					Use:            "update-warehouse [id] [name] [location] [capacity]",
					Short:          "Update Warehouse",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "location"}, {ProtoField: "capacity"}},
				},
				{
					RpcMethod:      "DeleteWarehouse",
					Use:            "delete-warehouse [id]",
					Short:          "Delete Warehouse",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
