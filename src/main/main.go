package main

import (
	"github.com/rbgayoivoye09/ssh-pcap/src/util/cmd"
	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
	"os"
)

func main() {
	if err := cmd.TrootCmd.Execute(); err != nil {
		Logger.Sugar().Error(err)
		os.Exit(1)
	}
}
