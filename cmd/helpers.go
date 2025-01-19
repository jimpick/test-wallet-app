package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	walletutils "github.com/glifio/go-wallet-utils"
)

func getAuthInstance(ctx context.Context, from string, msig string) (*walletutils.EthClientShim, *bind.TransactOpts, error) {
	if from == "" {
		return nil, nil, fmt.Errorf("key not found: %s", from)
	}

	ethClient, err := PoolsSDK.Extern().ConnectEthClient()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to instantiate eth client %s", err)
	}

	backends := []accounts.Backend{}
	backends = append(backends, KeyStore.KeyStore())
	manager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, backends...)

	var fromAddr interface{}
	var fromAddrType walletutils.KeyType

	fromAddr, fromAddrType, err = KeyStore.GetAddr(from)
	if err != nil {
		return nil, nil, err
	}

	if msig == "" {
		if fromAddrType == walletutils.KeyTypeEth {
			account := accounts.Account{Address: fromAddr.(common.Address)}
			wallet, err := manager.Find(account)
			if err != nil {
				return nil, nil, err
			}

			passphrase := ""
			err = KeyStore.KeyStore().Unlock(account, "")
			if err != nil {
				prompt := &survey.Password{Message: "Passphrase for account"}
				survey.AskOne(prompt, &passphrase)
				if passphrase == "" {
					return nil, nil, fmt.Errorf("Aborted")
				}
			}

			fmt.Println("Executing as ETH EOA")

			return walletutils.NewEthWalletTransactor(wallet, &account, passphrase, big.NewInt(chainID), ethClient)
		} else if fromAddrType == walletutils.KeyTypeFil {

			lapi, closer, err := PoolsSDK.Extern().ConnectLotusClient()
			if err != nil {
				return nil, nil, err
			}
			defer closer()

			var passphrase string
			var message = "Passphrase for account (or hit enter for no passphrase)"
			prompt := &survey.Password{Message: message}
			survey.AskOne(prompt, &passphrase)

			keyJSON, err := KeyStore.Export(from, passphrase)
			if err != nil {
				return nil, nil, err
			}

			pk, err := keystore.DecryptKey(keyJSON, passphrase)
			if err != nil {
				logFatal(err)
			}

			fmt.Println("Executing as FIL EOA")

			return walletutils.NewFilEoaWalletTransactor(ctx, lapi, ethClient, fromAddr.(address.Address), crypto.FromECDSA(pk.PrivateKey), passphrase)
		} else {
			return nil, nil, errors.New("unsupported from address type")
		}
	} else {
		msigAddr, msigAddrType, err := KeyStore.GetAddr(msig)
		if err != nil {
			return nil, nil, err
		}

		lapi, closer, err := PoolsSDK.Extern().ConnectLotusClient()
		if err != nil {
			return nil, nil, err
		}
		defer closer()

		var msigAddrFil address.Address
		var ok bool

		if msigAddrType == walletutils.KeyTypeFil {
			msigAddrFil, ok = msigAddr.(address.Address)
			if !ok {
				return nil, nil, errors.New("type assertion failed")
			}
		} else if msigAddrType == walletutils.KeyTypeEth {
			return nil, nil, errors.New("eth hex masked id addresses not supported")
		}

		if fromAddrType != walletutils.KeyTypeFil {
			return nil, nil, errors.New("from address must be a Filecoin account with keys in the wallet")
		}
		proposer, ok := fromAddr.(address.Address)
		if !ok {
			return nil, nil, errors.New("type assertion failed")
		}

		var passphrase string
		var message = "Passphrase for account (or hit enter for no passphrase)"
		prompt := &survey.Password{Message: message}
		survey.AskOne(prompt, &passphrase)

		keyJSON, err := KeyStore.Export(from, passphrase)
		if err != nil {
			return nil, nil, err
		}

		pk, err := keystore.DecryptKey(keyJSON, passphrase)
		if err != nil {
			logFatal(err)
		}
		proposerPrivateKey := crypto.FromECDSA(pk.PrivateKey)

		fmt.Println("Executing as FIL multisig proposal")

		return walletutils.NewFilMsigProposerWalletTransactor(ctx, lapi, ethClient, proposer, proposerPrivateKey, msigAddrFil)
	}
}
