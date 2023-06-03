package student

import (
	"context"
)

var Service = &studentService{}

// StudentServiceServer 实现学生服务
type studentService struct {
}

// GetStudent 实现 GetStudent RPC 方法
func (s *studentService) GetStudent(ct context.Context, req *StudentRequest) (*StudentResponse, error) {
	// 根据学生ID查询学生信息
	studentID := req.StudentId
	// TODO: 查询学生信息的逻辑，可以从数据库或其他数据源获取学生信息

	// 构建学生信息响应
	response := &StudentResponse{
		Student: &Student{
			Id:   studentID,
			Name: "John Doe",
			Age:  20,
		},
	}
	return response, nil
}

func (s *studentService) mustEmbedUnimplementedStudentServiceServer() {}
