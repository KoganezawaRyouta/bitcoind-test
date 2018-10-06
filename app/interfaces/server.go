package interfaces

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"fmt"

	pb "github.com/KoganezawaRyouta/bitcoind-test/app/pb"
)

type customerService struct {
	customers []*pb.Person
	m         sync.Mutex
}

func remove(persons []*pb.Person, person *pb.Person) []*pb.Person {
	result := []*pb.Person{}
	for _, v := range persons {
		if v.Name != person.Name {
			result = append(result, v)
		}
	}
	return result
}

func (cs *customerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {
	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
		cs.customers = remove(cs.customers, p)
	}
	return nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	fmt.Printf("result:%#v \n", p)
	return new(pb.ResponseType), nil
}

func NewServer() error {
	lis, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	server := grpc.NewServer()

	pb.RegisterCustomerServiceServer(server, new(customerService))
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	return nil
}
