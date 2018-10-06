package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/bitcoind-test/app/interfaces"
	"github.com/joho/godotenv"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var (
	server  string
	subject string
)

var subCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {

		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[Nuts Publisher] ")
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
			errsCh <- interfaces.PubBlockCount(server, subject)
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

var pubCmd = &cobra.Command{
	Use:   "publisher",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[Nuts Publisher] ")
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
			log.Printf(" Server : %s\n", server)
			errsCh <- interfaces.PubBlockCount(server, subject)
		}()
		log.Fatal("terminated", <-errsCh)

	},
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	subCmd.PersistentFlags().StringVar(&server, "s", os.Getenv("IP_FOR_NUTS_SERVER")+":"+os.Getenv("PORT_FOR_NUTS_SERVER"), "")
	subCmd.PersistentFlags().StringVar(&subject, "sub", "h", "")
	RootCmd.AddCommand(subCmd)

	pubCmd.PersistentFlags().StringVar(&server, "s", os.Getenv("IP_FOR_NUTS_SERVER")+":"+os.Getenv("PORT_FOR_NUTS_SERVER"), "")
	pubCmd.PersistentFlags().StringVar(&subject, "sub", "h", "")
	RootCmd.AddCommand(pubCmd)
}
