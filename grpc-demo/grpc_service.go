package main

import (
	"demos/grpc-demo/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	service.RegisterTestServiceServer(server, service.TestService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}
