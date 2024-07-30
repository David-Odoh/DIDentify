package identity

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"identity/testutil/sample"
	identitysimulation "identity/x/identity/simulation"
	"identity/x/identity/types"
)

// avoid unused import issue
var (
	_ = identitysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterIdentity = "op_weight_msg_register_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterIdentity int = 100

	opWeightMsgApproveIdentity = "op_weight_msg_approve_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveIdentity int = 100

	opWeightMsgCreateDid = "op_weight_msg_create_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterIdentity int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterIdentity, &weightMsgRegisterIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterIdentity = defaultWeightMsgRegisterIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterIdentity,
		identitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveIdentity int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveIdentity, &weightMsgApproveIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgApproveIdentity = defaultWeightMsgApproveIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveIdentity,
		identitysimulation.SimulateMsgApproveIdentity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDid int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDid, &weightMsgCreateDid, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDid = defaultWeightMsgCreateDid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDid,
		identitysimulation.SimulateMsgCreateDid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterIdentity,
			defaultWeightMsgRegisterIdentity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveIdentity,
			defaultWeightMsgApproveIdentity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgApproveIdentity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDid,
			defaultWeightMsgCreateDid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateDid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
