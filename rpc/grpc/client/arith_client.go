package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"mygrpc"
	"time"
)

var (
	addr = flag.String("addr", "localhost:1234", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mygrpc.NewArithClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	a := 10
	b := 2
	r, err := c.Multiply(ctx, &mygrpc.Param{A: int32(a), B: int32(b)})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("%d multiply %d = %d", a, b, r.Ret)

	q, err1 := c.Divide(ctx, &mygrpc.Param{A: int32(a), B: int32(b)})
	if err1 != nil {
		log.Fatalf("could not divide: %v", err1)
	}
	log.Printf("%d divide %d = %d, mod = %d", a, b, q.Quo, q.Rem)

	a, b = 20, 0
	q, err1 = c.Divide(ctx, &mygrpc.Param{A: int32(a), B: int32(b)})
	if err1 != nil {
		log.Fatalf("could not divide: %v", err1)
	}
	log.Printf("%d divide %d = %d, mod = %d", a, b, q.Quo, q.Rem)

	/*
		2022/06/20 23:24:40 10 multiply 2 = 20
		2022/06/20 23:24:40 10 divide 2 = 5, mod = 0
		2022/06/20 23:24:40 could not divide: rpc error: code = Unknown desc = divide by zero
	*/
}
