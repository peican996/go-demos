package main

import (
	"demos/grpc-demo/students/service/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册学生服务
	pb.RegisterQueryServiceServer(grpcServer, pb.QueryStudentsService)

	// 监听端口
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 启动 gRPC 服务器
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
