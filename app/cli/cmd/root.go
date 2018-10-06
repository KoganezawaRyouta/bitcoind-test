package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/bitcoind-test/app/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

var conf config.Config

func init() {
	cobra.OnInitialize(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		conf.LogFile = os.Getenv("LOG_FILE")
	})
}
