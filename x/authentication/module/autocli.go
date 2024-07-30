package authentication

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "identity/api/identity/authentication"
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
					RpcMethod:      "GenerateChallenge",
					Use:            "generate-challenge [did]",
					Short:          "Send a GenerateChallenge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
				},
				{
					RpcMethod:      "VerifyChallenge",
					Use:            "verify-challenge [did] [signed-challenge]",
					Short:          "Send a VerifyChallenge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "signedChallenge"}},
				},
				{
					RpcMethod:      "AuthenticateUser",
					Use:            "authenticate-user [did] [signed-challenge]",
					Short:          "Send a AuthenticateUser tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "signedChallenge"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
