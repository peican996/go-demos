package main

import (
	"context"
	"demos/grpc-demo/students/service/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	client := pb.NewQueryServiceClient(conn)

	// 构建学生信息查询请求
	request := &pb.QueryRequestInfo{Info: "select * from students where id = 2"}

	// 调用 GetStudent 方法获取学生信息
	response, err := client.GetStudents(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to get student: %v", err)
	}

	// 处理学生信息响应
	students := response.GetStudents()
	for _, student := range students {
		if student.GetId() != "" {
			log.Println("=====================================================")
			log.Printf("Student ID: %v", student.GetId())
			log.Printf("Student Name: %s", student.GetName())
			log.Printf("Student Age: %d", student.GetAge())
			log.Printf("Student Gender: %s", student.GetGender())
			log.Printf("Student StudentNumber: %s", student.GetStudentNumber())
			log.Printf("Student Grade: %s", student.GetGrade())
			log.Println("=====================================================")
		}
	}
}
