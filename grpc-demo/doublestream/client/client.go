package main

import (
	"context"
	"demos/grpc-demo/doublestream/client/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)
	stream, err := client.MyMethod(context.Background())
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}

	// 发送请求
	reqs := []*pb.Request{
		{Data: "Alice"},
		{Data: "Bob"},
		{Data: "Charlie"},
	}

	// 接收响应
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive response: %v", err)
			}
			// 处理响应
			log.Printf("Received response: %s", resp.Result)
		}
	}()

	for {
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed")
			}
		}
		time.Sleep(time.Second)
	}

	// 等待关闭
	//if err := stream.CloseSend(); err != nil {
	//	log.Fatalf("Failed to close stream: %v", err)
	//}

	// 阻塞等待
	//<-make(chan struct{})
}
