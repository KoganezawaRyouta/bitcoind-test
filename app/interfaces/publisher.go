package interfaces

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/KoganezawaRyouta/bitcoind-test/app/lib/btcrpc"
	"github.com/nats-io/go-nats"
)

func PubBlockCount(urls, subj string) error {
	nc, err := nats.Connect(urls)
	if err != nil {
		return err
	}
	defer nc.Close()

	client := btcrpc.NewClient()
	defer client.Shutdown()

	for {
		bc, err := client.GetBlockCount()
		if err != nil {
			log.Printf(" Process error : %v\n", err)
			return err
		}
		nc.Publish(subj, []byte(strconv.FormatInt(bc, 10)))
		log.Printf("Published [%s] : '%v'\n", subj, bc)
		nc.Flush()
		time.Sleep(1 * time.Second)
		if err := nc.LastError(); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Nats pub] ")
}
