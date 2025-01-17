/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	denoms "github.com/glifio/go-pools/util"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/spf13/cobra"
)

func printBalance(ctx context.Context, lapi *api.FullNodeStruct, name string, addr address.Address) {
	bal, err := lapi.WalletBalance(ctx, addr)
	if err != nil {
		fmt.Printf("%s balance: Error %v\n", name, err)
		return
	}
	balance := denoms.ToFIL(bal.Int)
	bf64, _ := balance.Float64()
	fmt.Printf("%s balance: %.02f FIL\n", name, bf64)
}

// newCmd represents the new command
var balCmd = &cobra.Command{
	Use:   "balance",
	Short: "Gets the balances associated with your accounts",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		accounts, err := KeyStore.List(false)
		if err != nil {
			logFatal(err)
		}

		lapi, closer, err := PoolsSDK.Extern().ConnectLotusClient()
		if err != nil {
			logFatalf("Failed to instantiate eth client %s", err)
		}
		defer closer()
		// loop through the list and convert all eth address to fil address types
		for i, acc := range accounts {
			if acc.KeyType == walletutils.KeyTypeEth {
				delegated, err := walletutils.DelegatedFromEthAddr(acc.Addr.(common.Address))
				if err != nil {
					logFatalf("failed to convert eth address to fil address: %s", err)
				}
				accounts[i].Addr = delegated
				accounts[i].KeyType = walletutils.KeyTypeFil
				printBalance(ctx, lapi, acc.Name, delegated)
			} else {
				printBalance(ctx, lapi, acc.Name, acc.Addr.(address.Address))
			}
		}
		fmt.Println()
	},
}

func init() {
	walletCmd.AddCommand(balCmd)
}
