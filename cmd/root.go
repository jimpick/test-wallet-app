/*
Copyright © 2023 Glif LTD
*/
package cmd

import (
	"context"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"

	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
	"github.com/glifio/go-pools/sdk"
	types "github.com/glifio/go-pools/types"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "calibnet"
var chainID int64 = constants.CalibnetChainID
var PoolsSDK types.PoolsSDK
var KeyStore *walletutils.KeyStorageShim

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "test-wallet-app",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var ExitCode int

func Exit(code int) {
	ExitCode = code
	runtime.Goexit()
}

func logFatal(arg interface{}) {
	log.Println(arg)
	Exit(1)
}

func logFatalf(format string, args ...interface{}) {
	log.Printf(format, args...)
	Exit(1)
}

func init() {
	viper.SetConfigName(cfgFile)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	initFlags()

	if err := viper.ReadInConfig(); err != nil {
		// log.Println("Failed to read config file:", err)
	}

	bindEnv()

	cobra.OnInitialize(initCLI)
}

// bindEnv binds environment variables for configuration values
func bindEnv() {
	// Daemon section
	viper.BindEnv("full_node")
	viper.BindEnv("full_node_token")
}

func initFlags() {
	rootCmd.PersistentFlags().BoolP("as-msig", "m", false, "Execute as a multisig transaction")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initCLI() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Define the config directory
	cfgDir := filepath.Join(home, ".test-wallet-app")
	if chainID == constants.CalibnetChainID {
		cfgDir = filepath.Join(home, ".test-wallet-app", "calibnet")
	} else if chainID != constants.MainnetChainID {
		logFatalf("Unsupported chain ID: %d", chainID)
	}

	// Check if the directory exists, and create it if it doesn't
	err = os.MkdirAll(cfgDir, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create config directory:", err)
	}

	// Initialize the keystore
	ks, err := walletutils.NewKeyStore(cfgDir)
	if err != nil {
		logFatalf("Failed to initialize keystore %s", err)
	}
	KeyStore = ks

	// defaults
	viper.SetDefault("lotus_url", "https://api.calibration.node.glif.io/rpc/v1")

	extern := deploy.Extern
	daemonURL := viper.GetString("lotus_url")
	daemonToken := viper.GetString("token")
	if daemonURL != "" {
		extern.LotusDialAddr = daemonURL
	}
	if daemonToken != "" {
		extern.LotusToken = daemonToken
	}

	sdk, err := sdk.New(context.Background(), big.NewInt(chainID), extern)
	if err != nil {
		logFatalf("Failed to initialize pools sdk %s", err)
	}
	PoolsSDK = sdk
}
