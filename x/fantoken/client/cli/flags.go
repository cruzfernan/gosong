package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagDenom     = "denom"
	FlagName      = "name"
	FlagMaxSupply = "max-supply"
	FlagMintable  = "mintable"
	FlagRecipient = "recipient"
	FlagAmount    = "amount"
)

var (
	FsIssueFanToken          = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateFanTokenMintable = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferFanTokenOwner  = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintFanToken           = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueFanToken.String(FlagDenom, "", "The token denom. Once created, it cannot be modified")
	FsIssueFanToken.String(FlagName, "", "The token name, e.g. IRIS Network")
	FsIssueFanToken.String(FlagMaxSupply, "", "The maximum supply of the token")
	FsIssueFanToken.Bool(FlagMintable, false, "Whether the token can be minted, default to false")

	FsUpdateFanTokenMintable.String(FlagName, "[do-not-modify]", "The token name, e.g. IRIS Network")
	FsUpdateFanTokenMintable.String(FlagMaxSupply, "", "The maximum supply of the token")
	FsUpdateFanTokenMintable.String(FlagMintable, "", "Whether the token can be minted, default to false")

	FsTransferFanTokenOwner.String(FlagRecipient, "", "The new owner")

	FsMintFanToken.String(FlagRecipient, "", "Address to which the token is to be minted")
	FsMintFanToken.String(FlagAmount, "", "Amount of the token to be minted")
}
