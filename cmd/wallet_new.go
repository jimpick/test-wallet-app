/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/spf13/cobra"
)

// walletNewCmd represents the create-account command
var walletNewCmd = &cobra.Command{
	Use:   "new [account-name]",
	Short: "Create a single named account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) == 1 {
			name = strings.ToLower(args[0])
		} else {
			name = "default"
		}

		typeFlag := walletutils.KeyTypeFromString(cmd.Flag("type").Value.String())
		if typeFlag != walletutils.KeyTypeEth && typeFlag != walletutils.KeyTypeFil {
			logFatalf("Invalid account type")
		}

		if err := walletutils.ValidateKeyName(name); err != nil {
			logFatal(err)
		}

		fmt.Printf("Creating %s account: %s\n", typeFlag, name)

		passphrase, envSet := os.LookupEnv("GLIF_WALLET_PASSPHRASE")
		if !envSet {
			prompt := &survey.Password{
				Message: "Please type a passphrase to encrypt your private key",
			}
			survey.AskOne(prompt, &passphrase)
			var confirmPassphrase string
			confirmPrompt := &survey.Password{
				Message: "Confirm passphrase",
			}
			survey.AskOne(confirmPrompt, &confirmPassphrase)
			if passphrase != confirmPassphrase {
				logFatal("Aborting. Passphrase confirmation did not match.")
			}
		}

		_, err := KeyStore.NewAccount(name, passphrase, typeFlag)
		if err != nil {
			logFatal(err)
		}

		addr, _, err := KeyStore.GetAddr(name)
		if err != nil {
			logFatal(err)
		}

		fmt.Printf("Account %s created successfully: %s\n", name, addr)
	},
}

func init() {
	walletCmd.AddCommand(walletNewCmd)
	walletNewCmd.Flags().String("type", "eth", "Type of account to create (eth or fil)")
}
