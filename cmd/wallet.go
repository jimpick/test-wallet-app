/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Commands for interacting with CLI wallets",
}

func init() {
	rootCmd.AddCommand(walletCmd)
}
