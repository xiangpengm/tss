package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

	"github.com/bnb-chain/tss/common"
	"github.com/google/uuid"
)

func init() {
	rootCmd.AddCommand(channelCmd)
}

var channelCmd = &cobra.Command{
	Use:              "channel",
	Short:            "generate a channel id for bootstrapping",
	TraverseChildren: false, // TODO: figure out how to disable parent's options
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("channel id: %s\n", fmt.Sprintf("%s", uuid.New().String()))
	},
}

func askChannelExpire() int {
	if viper.GetInt("channel_expire") > 0 {
		return viper.GetInt("channel_expire")
	}

	reader := bufio.NewReader(os.Stdin)
	expire, err := common.GetInt("please set expire time in minutes, (default: 30): ", 30, reader)
	if err != nil {
		common.Panic(err)
	}
	if expire <= 0 {
		common.Panic(fmt.Errorf("expire time should not be zero or negative value"))
	}
	return expire
}
