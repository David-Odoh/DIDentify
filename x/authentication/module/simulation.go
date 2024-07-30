package authentication

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"identity/testutil/sample"
	authenticationsimulation "identity/x/authentication/simulation"
	"identity/x/authentication/types"
)

// avoid unused import issue
var (
	_ = authenticationsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgGenerateChallenge = "op_weight_msg_generate_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenerateChallenge int = 100

	opWeightMsgVerifyChallenge = "op_weight_msg_verify_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyChallenge int = 100

	opWeightMsgAuthenticateUser = "op_weight_msg_authenticate_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuthenticateUser int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	authenticationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&authenticationGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgGenerateChallenge int
	simState.AppParams.GetOrGenerate(opWeightMsgGenerateChallenge, &weightMsgGenerateChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgGenerateChallenge = defaultWeightMsgGenerateChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenerateChallenge,
		authenticationsimulation.SimulateMsgGenerateChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyChallenge int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyChallenge, &weightMsgVerifyChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyChallenge = defaultWeightMsgVerifyChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyChallenge,
		authenticationsimulation.SimulateMsgVerifyChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAuthenticateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgAuthenticateUser, &weightMsgAuthenticateUser, nil,
		func(_ *rand.Rand) {
			weightMsgAuthenticateUser = defaultWeightMsgAuthenticateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAuthenticateUser,
		authenticationsimulation.SimulateMsgAuthenticateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgGenerateChallenge,
			defaultWeightMsgGenerateChallenge,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				authenticationsimulation.SimulateMsgGenerateChallenge(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyChallenge,
			defaultWeightMsgVerifyChallenge,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				authenticationsimulation.SimulateMsgVerifyChallenge(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAuthenticateUser,
			defaultWeightMsgAuthenticateUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				authenticationsimulation.SimulateMsgAuthenticateUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
