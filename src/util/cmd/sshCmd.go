package cmd

import (
	"github.com/rbgayoivoye09/ssh-pcap/src/util/config"
	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
	"github.com/rbgayoivoye09/ssh-pcap/src/util/ssh"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into a remote server",
	Run: func(cmd *cobra.Command, args []string) {

		c := config.GetConfig(inputConfigFilePath)

		if len(c.SSHConfig) == 0 {

		} else {
			for k, v := range c.SSHConfig {
				Logger.Sugar().Infof("%s: %s", k, v)
				if str, err := ssh.ExecuteRemoteCommand(v.Host, v.Port, v.Username, v.Password, v.PcapCmd); err != nil {
					Logger.Sugar().Error(err)
				} else {
					Logger.Sugar().Info(str)
				}
			}
		}

	},
}
