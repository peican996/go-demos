package main

import (
	"demos/grpc-demo/twowaystreaming/service/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 创建监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}
	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 注册 FileTransferService 服务
	pb.RegisterFileTransferServiceServer(server, pb.FileTransferServer)

	// 启动服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
