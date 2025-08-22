package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bnb-chain/tss/client"
	"github.com/bnb-chain/tss/common"
)

func init() {
	rootCmd.AddCommand(signCmd)
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "sign a transaction",
	Long:  "sign a transaction using local share, signers will be prompted to fill in",
	PreRun: func(cmd *cobra.Command, args []string) {
		vault := askVault()
		passphrase := askPassphrase()
		if err := common.ReadConfigFromHome(viper.GetViper(), false, viper.GetString(flagHome), vault, passphrase); err != nil {
			common.Panic(err)
		}
		initLogLevel(common.TssCfg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		setChannelId()
		setChannelPasswd()
		setMessage()

		c := client.NewTssClient(&common.TssCfg, client.SignMode, false)
		c.Start()
	},
}

func setMessage() {
	message := viper.GetString("message")
	if message == "" {
		message = "0" // Default value if not provided
	}
	common.TssCfg.Message = message
}
