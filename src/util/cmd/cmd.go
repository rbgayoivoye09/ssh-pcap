package cmd

import (
	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
	"github.com/spf13/cobra"
)

var inputConfigFilePath string

var TrootCmd = &cobra.Command{
	Use: "ssh-pcap", Short: "ssh pcap commands",
	Run: func(cmd *cobra.Command, args []string) {
		if s, err := cmd.Flags().GetString("config"); err != nil {
			Logger.Sugar().Error(err.Error())
		} else {
			inputConfigFilePath = s
		}
		Logger.Sugar().Infof("config file path: %s", inputConfigFilePath)

	},
}

func init() {
	TrootCmd.AddCommand(sshCmd)
	TrootCmd.PersistentFlags().StringVarP(&inputConfigFilePath, "config", "c", "", "custom config file path")

}
