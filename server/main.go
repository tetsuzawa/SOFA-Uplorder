package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/tetsuzawa/SOFA-Uplorder/service"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	service.NewUploadServer(server)
	fmt.Println("start server...")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
