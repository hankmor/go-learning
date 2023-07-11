package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo"
	"log"
	"net"
)

var (
	port = flag.Int("port", 1234, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	mygrpc.UnimplementedArithServer
}

// Multiply implements
func (s *server) Multiply(ctx context.Context, in *mygrpc.Param) (*mygrpc.Int, error) {
	log.Printf("Received: %v, %v", in.GetA(), in.GetB())
	return &mygrpc.Int{Ret: (in.A * in.B)}, nil
}

func (s *server) Divide(ctx context.Context, in *mygrpc.Param) (*mygrpc.Quotient, error) {
	log.Printf("Received: %v, %v", in.GetA(), in.GetB())
	if in.B == 0 {
		return nil, errors.New("divide by zero")
	}
	return &mygrpc.Quotient{Quo: in.A / in.B, Rem: in.A % in.B}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mygrpc.RegisterArithServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
