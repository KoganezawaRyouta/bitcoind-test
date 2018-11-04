package btcrpc

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
)

func NewClient() *rpcclient.Client {
	log.Printf(" RPC_USER : %v\n", os.Getenv("RPC_USER"))
	log.Printf(" RPC_PASSWORD : %v\n", os.Getenv("RPC_PASSWORD"))
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:18332",
		User:         os.Getenv("RPC_USER"),
		Pass:         os.Getenv("RPC_PASSWORD"),
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
