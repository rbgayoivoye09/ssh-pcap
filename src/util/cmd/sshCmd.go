package cmd

import (
	"github.com/rbgayoivoye09/ssh-pcap/src/util/config"
	"github.com/rbgayoivoye09/ssh-pcap/src/util/flag"
	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
	"github.com/rbgayoivoye09/ssh-pcap/src/util/ssh"
	"github.com/spf13/cobra"
)

func init() {
	sshCmd.Flags().BoolP(flag.PCAP_FLAG, "p", false, "pcap command")
	sshCmd.Flags().BoolP(flag.STOP_FLAG, "s", false, "stop pcap command")
	sshCmd.Flags().BoolP(flag.DOWN_FLAG, "d", false, "download pcap command")
	sshCmd.Flags().Bool(flag.SHOW_FLAG, false, "show pcap tasks")
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into a remote server",
	Run: func(cmd *cobra.Command, args []string) {

		c := config.GetConfig(inputConfigFilePath)

		if len(c.SSHConfig) == 0 {
			Logger.Sugar().Error("No SSH servers found in config file")
		} else {
			for k, v := range c.SSHConfig {
				Logger.Sugar().Infof("%s: %s", k, v)
				if b, err := cmd.Flags().GetBool(flag.PCAP_FLAG); err == nil {
					if b {
						Logger.Sugar().Info("pcap: ", v.PcapCmd)

						if str, err := ssh.ExecuteRemoteCommand(v.Host, v.Port, v.Username, v.Password, flag.PCAP_FLAG, v.PcapCmd, c.LocalFilePath); err != nil {
							Logger.Sugar().Error(err)
						} else {
							Logger.Sugar().Info(str)
						}
					}
				} else {

				}

				if b, err := cmd.Flags().GetBool(flag.STOP_FLAG); err == nil {
					if b {
						Logger.Sugar().Info("stop: ", v.StopCmd)

						if str, err := ssh.ExecuteRemoteCommand(v.Host, v.Port, v.Username, v.Password, flag.STOP_FLAG, v.StopCmd, c.LocalFilePath); err != nil {
							Logger.Sugar().Error(err)
						} else {
							Logger.Sugar().Info(str)
						}
					}
				} else {

				}
				if b, err := cmd.Flags().GetBool(flag.DOWN_FLAG); err == nil {
					if b {
						Logger.Sugar().Info("download: ", v.DownCmd)

						if str, err := ssh.ExecuteRemoteCommand(v.Host, v.Port, v.Username, v.Password, flag.DOWN_FLAG, v.DownCmd, c.LocalFilePath); err != nil {
							Logger.Sugar().Error(err)
						} else {
							Logger.Sugar().Info(str)
						}
					}
				} else {

				}

				if b, err := cmd.Flags().GetBool(flag.SHOW_FLAG); err == nil {
					if b {
						Logger.Sugar().Info("show: ", v.ShowCmd)

						if str, err := ssh.ExecuteRemoteCommand(v.Host, v.Port, v.Username, v.Password, flag.SHOW_FLAG, v.ShowCmd, c.LocalFilePath); err != nil {
							Logger.Sugar().Error(err)
						} else {
							Logger.Sugar().Info(str)
						}
					}

				} else {

				}

			}
		}

	},
}
