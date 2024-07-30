package identity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "identity/api/identity/identity"
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
					RpcMethod: "IdentityAll",
					Use:       "list-identity",
					Short:     "List all identity",
				},
				{
					RpcMethod:      "Identity",
					Use:            "show-identity [id]",
					Short:          "Shows a identity by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "SearchIdentity",
					Use:            "search-identity [national-id]",
					Short:          "Query searchIdentity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nationalId"}},
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
					RpcMethod:      "RegisterIdentity",
					Use:            "register-identity [name] [birthdate] [national-id]",
					Short:          "Send a registerIdentity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "birthdate"}, {ProtoField: "nationalId"}},
				},
				{
					RpcMethod:      "ApproveIdentity",
					Use:            "approve-identity [id] [name] [birthdate] [national-id]",
					Short:          "Send a approveIdentity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "birthdate"}, {ProtoField: "nationalId"}},
				},
				{
					RpcMethod:      "CreateDid",
					Use:            "create-did",
					Short:          "Send a createDID tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletIdentifier"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
