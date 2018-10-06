package interfaces

import (
	"log"
	"os"
	"runtime"

	"github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func NewSubscriber(urls string, subj string) error {
	nc, err := nats.Connect(urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", urls)
		return err
	}
	log.Printf("Connected on [%s]\n", urls)

	i := 0
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Listening on [%s]\n", subj)
	runtime.Goexit()
	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Nats sub] ")
}
