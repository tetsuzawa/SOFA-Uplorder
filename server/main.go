package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/tetsuzawa/SOFA-Uplorder/service"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	service.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
