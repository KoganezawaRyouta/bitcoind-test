package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/bitcoind-test/app/interfaces"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var btcCmd = &cobra.Command{
	Use:   "btcCmd",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[GRPC Server] ")
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			log.Printf("start")
			log.Printf(" PID          : %d\n", pidInfo.Pid())
			log.Printf(" PPID         : %d\n", pidInfo.PPid())
			log.Printf(" Process name : %s\n", pidInfo.Executable())
			pp, _ := ps.FindProcess(pidInfo.PPid())
			if pp != nil {
				log.Printf(" Parent process name : %s\n", pp.Executable())
			}
			errsCh <- interfaces.GetInfo()
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

func init() {
	RootCmd.AddCommand(btcCmd)
}
