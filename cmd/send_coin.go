/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/jimpick/test-wallet-app/abigen"
	"github.com/spf13/cobra"
)

var SIMPLECOIN = common.HexToAddress("0x784B1802F006D6Ac9e2F6758BcA882060bfc1eD5")

var sendCoinCmd = &cobra.Command{
	Use:   "send-coin <from> <to> <number-of-coins>",
	Short: "Sends simple coins from one account to another account",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		from := args[0]
		msig := ""

		to, keyType, err := KeyStore.GetAddr(args[1])
		if err != nil {
			log.Fatal(err)
		}

		if keyType != walletutils.KeyTypeEth {
			log.Fatal("To address must be eth")
		}
		toEth := to.(common.Address)

		amount, err := strconv.ParseUint(args[2], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		gasMultiply, err := cmd.Flags().GetFloat64("gas-multiply")
		if err != nil {
			log.Fatal(err)
		}

		nonce, err := cmd.Flags().GetUint64("nonce")
		if err != nil {
			log.Fatal(err)
		}

		// from, msig := getFromAndMsigFromFlags(cmd)

		ethClientShim, auth, err := getAuthInstance(cmd.Context(), from, msig)
		if err != nil {
			logFatal(err)
		}

		simpleCoinCaller, err := abigen.NewSimpleCoinTransactor(SIMPLECOIN, ethClientShim)
		if err != nil {
			logFatal(err)
		}

		tipCap, err := ethClientShim.SuggestGasTipCap(ctx)
		if err != nil {
			logFatal(err)
		}
		log.Printf("Suggested GasTipCap: %v\n", tipCap)
		tipCapFloat, _ := tipCap.Float64()
		scaledTipCap := tipCapFloat * gasMultiply
		auth.GasTipCap = big.NewInt(int64(scaledTipCap))
		log.Printf("Scaled GasTipCap: %v\n", auth.GasTipCap)

		if nonce > 0 {
			auth.Nonce = big.NewInt(int64(nonce))
		}

		tx, err := simpleCoinCaller.SendCoin(auth, toEth, big.NewInt(int64(amount)))
		if err != nil {
			logFatal(err)
		}

		outerTx := ethClientShim.GetOuterTxHash(tx.Hash())
		log.Println("Transaction hash: ", outerTx.String())
	},
}

func init() {
	rootCmd.AddCommand(sendCoinCmd)
	sendCoinCmd.Flags().Float64("gas-multiply", 1.0, "Multiply the default gas premium by this amount")
	sendCoinCmd.Flags().Uint64("nonce", 0, "Specify nonce (for replacing transactions)")
}
