/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/spf13/cobra"
)

// walletImportRaw represents the import-account command
var walletImport = &cobra.Command{
	Use:   "import [account-name] [account-private-key] [flags]",
	Short: "Import a single private key account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		overwrite, err := cmd.Flags().GetBool("overwrite")
		if err != nil {
			logFatal(err)
		}

		typeFlag := walletutils.KeyTypeFromString(cmd.Flag("type").Value.String())
		if typeFlag != walletutils.KeyTypeEth && typeFlag != walletutils.KeyTypeFil {
			logFatal("Invalid account type")
		}

		keyFormat := cmd.Flag("format").Value.String()
		if keyFormat != "raw-hex" && keyFormat != "encrypted-hex-json" && keyFormat != "raw-b64" {
			logFatal("Invalid format")
		}

		if keyFormat == "raw-b64" {
			// convert base64 to hex and reset the keyFormat
			decoded, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				logFatal(err)
			}
			args[1] = hex.EncodeToString(decoded)
			keyFormat = "raw-hex"
		}

		name := strings.ToLower(args[0])

		var e *walletutils.ErrKeyNotFound
		_, _, err = KeyStore.GetAddr(name)
		if !errors.As(err, &e) && !overwrite {
			logFatal("Account already exists")
		} else if !errors.As(err, &e) {
			// if we are overwriting, rename the account for safety
			rename := fmt.Sprintf("%s-replaced-%s", name, time.Now().Format(time.RFC3339))
			log.Printf("Warning: account '%s' already exists, renaming to '%s' and overriding with new '%s' key\n", name, rename, name)
			if err := KeyStore.Rename(name, rename); err != nil {
				logFatal(err)
			}
		}

		var passphrase string
		survey.AskOne(&survey.Password{Message: "Passphrase for account (or hit enter for no passphrase)"}, &passphrase)

		if err := KeyStore.Import(name, args[1], passphrase, typeFlag, keyFormat == "encrypted-hex-json"); err != nil {
			logFatal(err)
		}

		addr, keytype, err := KeyStore.GetAddr(name)
		if err != nil {
			logFatal(err)
		}

		fmt.Printf("Successfully imported %s account: %s\n", keytype, addr)
	},
}

func init() {
	walletCmd.AddCommand(walletImport)
	walletImport.Flags().Bool("overwrite", false, "overwrite an existing account with the same name")
	walletImport.Flags().String("format", "raw-hex", "What form the private key is in (raw-hex, encrypted-hex-json, raw-b64)")
	walletImport.Flags().String("type", "eth", "Type of account to create (eth or fil)")
}
