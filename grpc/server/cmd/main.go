package main

import (
	"flag"
	"google.golang.org/grpc"
	"my-go/grpc/server/service"

	"my-go/grpc/client"

	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":7999")
	if err != nil {
		panic(err)
	}

	svr := grpc.NewServer()
	client.RegisterSearchServer(svr, &service.Service{})
	_ = svr.Serve(lis)
}
