package proto

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	request []*GetRequest
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func main() {
	//server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &server{})
	srv.Serve(lis)

	//client
	conn, err := grpc.Dial("user:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.GetRequest{
		Keyword: "params request",
	}
	resp, err := client.CreateUser(context.Background(), request)
	// To do something with resp from instance server response
}
