/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var changePassphraseCmd = &cobra.Command{
	Use:   "change-passphrase <account or address>",
	Short: "Change the passphrase for an encrypted key in the keystore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oldPassphrase := ""
		survey.AskOne(&survey.Password{Message: "Old passphrase"}, &oldPassphrase)

		newPassphrase := ""
		survey.AskOne(&survey.Password{Message: "New passphrase"}, &newPassphrase)

		var confirmPassphrase string
		survey.AskOne(&survey.Password{Message: "Confirm passphrase"}, &confirmPassphrase)
		if newPassphrase != confirmPassphrase {
			logFatal("Aborting. Passphrase confirmation did not match.")
		}

		if err := KeyStore.ChangePassphrase(args[0], oldPassphrase, newPassphrase); err != nil {
			logFatal(err)
		}

		log.Printf("Passphrase for account %s successfully changed\n", args[0])
	},
}

func init() {
	walletCmd.AddCommand(changePassphraseCmd)
}
