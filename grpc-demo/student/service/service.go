package main

import (
	"demos/grpc-demo/student/service/student"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册学生服务
	student.RegisterStudentServiceServer(grpcServer, student.Service)

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
