package auth

import "github.com/cosmos/cosmos-sdk/x/auth"

const (
	ModuleName        = auth.ModuleName
	StoreKey          = auth.StoreKey
	FeeCollectorName  = auth.FeeCollectorName
	DefaultParamspace = auth.DefaultParamspace
)

var (
	// functions aliases
	NewAnteHandler                    = auth.NewAnteHandler
	DefaultSigVerificationGasConsumer = auth.DefaultSigVerificationGasConsumer
	NewAccountKeeper                  = auth.NewAccountKeeper
	ProtoBaseAccount                  = auth.ProtoBaseAccount
	DefaultTxDecoder                  = auth.DefaultTxDecoder
	NewCosmosAppModule                = auth.NewAppModule

	// variable aliases
	CosmosModuleCdc = auth.ModuleCdc
)

type (
	AccountKeeper        = auth.AccountKeeper
	CosmosAppModule      = auth.AppModule
	CosmosAppModuleBasic = auth.AppModuleBasic
)
