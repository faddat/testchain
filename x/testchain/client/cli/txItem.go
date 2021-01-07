package cli

import (
	"bufio"
    "strconv"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/faddat/testchain/x/testchain/types"
)

func GetCmdCreateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-item [description] [price] [image]",
		Short: "Creates a new item",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDescription := string(args[0] )
			argsPrice, _ := strconv.ParseInt(args[1] , 10, 64)
			argsImage := string(args[2] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateItem(cliCtx.GetFromAddress(), string(argsDescription), int32(argsPrice), string(argsImage))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-item [id]  [description] [price] [image]",
		Short: "Set a new item",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsDescription := string(args[1])
			argsPrice, _ := strconv.ParseInt(args[2], 10, 64)
			argsImage := string(args[3])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetItem(cliCtx.GetFromAddress(), id, string(argsDescription), int32(argsPrice), string(argsImage))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-item [id]",
		Short: "Delete a new item by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteItem(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
