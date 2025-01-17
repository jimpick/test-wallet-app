/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the addresses associated with your accounts",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := KeyStore.List(cmd.Flags().Changed("include-read-only"))
		if err != nil {
			logFatal(err)
		}

		if len(entries) == 0 {
			fmt.Println("No accounts found")
			return
		}

		for _, entry := range entries {
			fmt.Printf("%s: %s\n", entry.Name, entry.Addr)
		}
		fmt.Println()
	},
}

func init() {
	walletCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("include-read-only", false, "Include read-only wallet accounts in the list")
}
