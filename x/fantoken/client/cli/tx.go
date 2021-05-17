package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bitsongofficial/bitsong/x/fantoken/types"
)

// NewTxCmd returns the transaction commands for the token module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Token transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdIssueFanToken(),
		GetCmdEditFanToken(),
		GetCmdMintFanToken(),
		GetCmdBurnFanToken(),
		GetCmdTransferFanTokenOwner(),
	)

	return txCmd
}

// GetCmdIssueFanToken implements the issue fan token command
func GetCmdIssueFanToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "issue",
		Long: "Issue a new fan token.",
		Example: fmt.Sprintf(
			"$ %s tx token issue "+
				"--name=\"Kitty Token\" "+
				"--symbol=\"kitty\" "+
				"--max-supply=\"1000000000000\" "+
				"--description=\"Kitty Token\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()
			symbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return err
			}
			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}
			maxSupplyStr, err := cmd.Flags().GetString(FlagMaxSupply)
			if err != nil {
				return err
			}
			maxSupply, ok := sdk.NewIntFromString(maxSupplyStr)
			if !ok {
				return fmt.Errorf("failed to parse max supply: %s", maxSupplyStr)
			}
			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			msg := &types.MsgIssueFanToken{
				Symbol:      symbol,
				Name:        name,
				MaxSupply:   maxSupply,
				Description: description,
				Owner:       owner.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsIssueFanToken)
	_ = cmd.MarkFlagRequired(FlagSymbol)
	_ = cmd.MarkFlagRequired(FlagName)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditFanToken implements the update fan token mintable command
func GetCmdEditFanToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update [symbol]",
		Long: "Update an existing fan token mintable.",
		Example: fmt.Sprintf(
			"$ %s tx token edit <symbol> "+
				"--mintable=true "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			rawMintable, err := cmd.Flags().GetString(FlagMintable)
			if err != nil {
				return err
			}
			mintable, err := strconv.ParseBool(rawMintable)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditFanToken(args[0], mintable, owner)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsEditFanToken)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdMintFanToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mint [symbol]",
		Long: "Mint tokens to a specified address.",
		Example: fmt.Sprintf(
			"$ %s tx token mint <symbol> "+
				"--recipient=<recipient>"+
				"--amount=<amount> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			amountStr, err := cmd.Flags().GetString(FlagAmount)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(amountStr)
			if !ok {
				return fmt.Errorf("invalid amount %s", amount)
			}

			addr, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}
			if len(addr) > 0 {
				if _, err = sdk.AccAddressFromBech32(addr); err != nil {
					return err
				}
			}

			msg := types.NewMsgMintFanToken(
				addr, strings.TrimSpace(args[0]), owner, amount,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsMintFanToken)
	_ = cmd.MarkFlagRequired(FlagAmount)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdBurnFanToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "burn [symbol]",
		Long: "Burn fan tokens.",
		Example: fmt.Sprintf(
			"$ %s tx token burn <symbol> "+
				"--amount=<amount> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			amountStr, err := cmd.Flags().GetString(FlagAmount)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(amountStr)
			if !ok {
				return fmt.Errorf("invalid amount %s", amount)
			}

			addr, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}
			if len(addr) > 0 {
				if _, err = sdk.AccAddressFromBech32(addr); err != nil {
					return err
				}
			}

			msg := types.NewMsgBurnFanToken(
				strings.TrimSpace(args[0]), owner, amount,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsMintFanToken)
	_ = cmd.MarkFlagRequired(FlagAmount)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTransferFanTokenOwner implements the transfer fan token owner command
func GetCmdTransferFanTokenOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transfer [symbol]",
		Long: "Transfer the owner of a token to a new owner.",
		Example: fmt.Sprintf(
			"$ %s tx token transfer <symbol> "+
				"--to=<to> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			toAddr, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(toAddr); err != nil {
				return err
			}

			msg := types.NewMsgTransferFanTokenOwner(strings.TrimSpace(args[0]), owner, toAddr)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsTransferFanTokenOwner)
	_ = cmd.MarkFlagRequired(FlagRecipient)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
