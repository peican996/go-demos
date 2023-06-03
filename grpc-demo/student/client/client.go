package main

import (
	"context"
	"demos/grpc-demo/student/service/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 连接 gRPC 服务器
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 创建学生服务客户端
	client := student.NewStudentServiceClient(conn)

	// 构建学生信息查询请求
	request := &student.StudentRequest{
		StudentId: "12345",
	}

	// 调用 GetStudent 方法获取学生信息
	response, err := client.GetStudent(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to get student: %v", err)
	}

	// 处理学生信息响应
	student := response.GetStudent()
	log.Printf("Student ID: %s", student.GetId())
	log.Printf("Student Name: %s", student.GetName())
	log.Printf("Student Age: %d", student.GetAge())
}
