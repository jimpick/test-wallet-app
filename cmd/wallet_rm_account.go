/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// walletRemoveCmd represents the remove command
var walletRemoveCmd = &cobra.Command{
	Use:   "remove [account-name]",
	Short: "Remove an account and its private key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reallyDo, err := cmd.Flags().GetBool("really-do-it")
		if err != nil {
			logFatal(err)
		}

		if !reallyDo {
			logFatal("DANGEROUS COMMAND - are you really trying to export a raw private key from your wallet? If so, you must pass --really-do-it to complete the export")
		}

		name := strings.ToLower(args[0])
		var passphrase string
		var message = "Passphrase for account (or hit enter for no passphrase)"
		prompt := &survey.Password{Message: message}
		survey.AskOne(prompt, &passphrase)

		if err := KeyStore.Delete(name, passphrase); err != nil {
			logFatal(err)
		}

		log.Printf("Account %s removed successfully\n", name)
	},
}

func init() {
	walletCmd.AddCommand(walletRemoveCmd)
	walletRemoveCmd.Flags().Bool("really-do-it", false, "really remove the account")
}
