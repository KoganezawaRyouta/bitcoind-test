package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/bitcoind-test/app/interfaces"
	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var clientListCmd = &cobra.Command{
	Use:   "grpc_client_list",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[GRPC Client List] ")
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
			errsCh <- interfaces.NewList()
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

var clientAddCmd = &cobra.Command{
	Use:   "grpc_client_add",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[GRPC Client Add] ")
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
			errsCh <- interfaces.NewAdd()
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

func init() {
	RootCmd.AddCommand(clientListCmd)
	RootCmd.AddCommand(clientAddCmd)
}
