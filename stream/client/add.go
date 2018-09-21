package client

import (
	"context"
	"log"

	pb "github.com/KoganezawaRyouta/bitcoind-test/stream/pb"
	"google.golang.org/grpc"
)

//
//var (
//	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
//	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
//	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
//)

func add(name string, age int) error {
	conn, err := grpc.Dial("grpc-server:2222", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
		Age:  int32(age),
	}
	_, err = client.AddPerson(context.Background(), person)
	return err
}

func NewAdd() error {
	err := add("a", 100)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	err = add("b", 200)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	err = add("c", 300)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	err = add("d", 400)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	err = add("e", 500)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	return nil
}
