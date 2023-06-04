package main

import (
	"context"
	"demos/grpc-demo/twowaystreaming/service/pb"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	// 创建到服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到服务器: %v", err)
	}
	defer conn.Close()
	// 创建 FileTransferService 的客户端
	client := pb.NewFileTransferServiceClient(conn)
	// 创建双向流式传输
	stream, err := client.TransferFile(context.Background())
	if err != nil {
		log.Fatalf("无法创建流: %v", err)
	}
	// 发送文件名
	err = stream.Send(&pb.FileRequest{Filename: "grpc-demo/twowaystreaming/file/service/datagrip-2022.3.3.exe"})
	if err != nil {
		log.Fatalf("发送文件名失败: %v", err)
	}
	// 关闭发送
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("关闭发送失败: %v", err)
	}
	file, err := os.Create("grpc-demo/twowaystreaming/file/client/datagrip-2022.3.3.exe")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	timestampBegin := time.Now().Unix()
	fmt.Println("当前时间戳（秒级）：", timestampBegin)
	// 接收服务器返回的文件数据
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("接收文件数据失败: %v", err)
		}
		// 处理接收到的文件数据
		fileData := resp.GetFiledata()
		_, err = file.Write(fileData)
		if err != nil {
			log.Fatal(err)
		}
	}

	timestampEnd := time.Now().Unix()
	fmt.Println("传输耗时：", timestampEnd-timestampBegin)
}
