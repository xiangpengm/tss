package cmd

import (
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bnb-chain/tss/client"
	"github.com/bnb-chain/tss/common"
)

func init() {
	rootCmd.AddCommand(describeCmd)
}

// fmt.Printf is deliberately used in this command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "show config and address of a tss vault",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		vault := askVault()
		passphrase := askPassphrase()
		if err := common.ReadConfigFromHome(viper.GetViper(), false, viper.GetString(flagHome), vault, passphrase); err != nil {
			common.Panic(err)
		}
		initLogLevel(common.TssCfg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		pubKey, err := common.LoadEcdsaPubkey(viper.GetString(flagHome), viper.GetString(flagVault), common.TssCfg.Password)
		if err != nil {
			fmt.Printf("cannot load public key, maybe not keygen yet: %v", err)
		}
		if pubKey != nil {
			addr, err := client.GetAddress(*pubKey, viper.GetString(flagPrefix))
			if err != nil {
				common.Panic(err)
			}
			compressed := elliptic.MarshalCompressed(pubKey.Curve, pubKey.X, pubKey.Y)
			fmt.Printf("pubkey of this vault: %s\n", hex.EncodeToString(compressed))
			fmt.Printf("address of this vault: %s\n", addr)
		}
		cfg, err := json.MarshalIndent(common.TssCfg, "", "\t")
		if err != nil {
			common.Panic(err)
		}
		fmt.Printf("config of this vault:\n%s\n", string(cfg))
	},
}
