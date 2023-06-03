package main

import (
	"context"
	"demos/grpc-demo/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8002端口
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 2. 调用****.pb.go中的New****方法
	grpcServiceClient := service.NewTestServiceClient(conn)

	// 3. 直接像调用本地方法一样调用Service中方法
	resp, err := grpcServiceClient.ServiceMethodInvoke(context.Background(), &service.ServiceParam{ServiceParam1: 100, ServiceParam2: 10})
	if err != nil {
		log.Fatal("Error calling gRPC method: ", err)
	}

	fmt.Printf("获取取服务器中的参数1: %v,获取取服务器中的参数2: %v\n", resp.GetClientParam1(), resp.GetClientParam2())
}
