/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"log"

	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/spf13/cobra"
)

// labelAccountCmd represents the label-account command
var labelAccountCmd = &cobra.Command{
	Use:   "label-account <name> <address>",
	Short: "Label an account with a human readable name",
	Long:  "Labeling an account creates a read-only alias for an account's address.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// if its a filecoin address, we check to make sure its valid
		keytype := walletutils.KeyTypeFromAddr(args[1])

		if keytype == walletutils.KeyTypeUnknown {
			logFatalf("Invalid address")
		}

		err := KeyStore.NewReadOnlyAccount(args[0], args[1])
		if err != nil {
			logFatal(err)
		}

		log.Printf("Successfully added new read-only account %s to wallet - %s\n", args[0], args[1])
	},
}

func init() {
	walletCmd.AddCommand(labelAccountCmd)
}
