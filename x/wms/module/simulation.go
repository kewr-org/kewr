package wms

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kewr/testutil/sample"
	wmssimulation "kewr/x/wms/simulation"
	"kewr/x/wms/types"
)

// avoid unused import issue
var (
	_ = wmssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWarehouse int = 100

	opWeightMsgUpdateWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWarehouse int = 100

	opWeightMsgDeleteWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWarehouse int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	wmsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		WarehouseList: []types.Warehouse{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		WarehouseCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&wmsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateWarehouse int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateWarehouse, &weightMsgCreateWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWarehouse = defaultWeightMsgCreateWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWarehouse,
		wmssimulation.SimulateMsgCreateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWarehouse int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateWarehouse, &weightMsgUpdateWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWarehouse = defaultWeightMsgUpdateWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWarehouse,
		wmssimulation.SimulateMsgUpdateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWarehouse int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteWarehouse, &weightMsgDeleteWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWarehouse = defaultWeightMsgDeleteWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWarehouse,
		wmssimulation.SimulateMsgDeleteWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateWarehouse,
			defaultWeightMsgCreateWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				wmssimulation.SimulateMsgCreateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateWarehouse,
			defaultWeightMsgUpdateWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				wmssimulation.SimulateMsgUpdateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteWarehouse,
			defaultWeightMsgDeleteWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				wmssimulation.SimulateMsgDeleteWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
