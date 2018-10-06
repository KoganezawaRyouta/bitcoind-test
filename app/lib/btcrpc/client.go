package btcrpc

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
)

func NewClient() *rpcclient.Client {
	connCfg := &rpcclient.ConnConfig{
		Host:         os.Getenv("RPC_BIND") + ":18332",
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
