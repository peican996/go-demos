package pb

import (
	"context"
	"demos/grpc-demo/students/database/database"
	"demos/grpc-demo/students/database/utils"
	"fmt"
	"log"
)

var QueryStudentsService = &queryService{}

type queryService struct {
}

func (qs *queryService) GetStudents(ct context.Context, queryInfo *QueryRequestInfo) (*QueryResponse, error) {
	students := getStudents(queryInfo)
	reponse := &QueryResponse{Students: students}
	return reponse, nil
}

func (qs *queryService) mustEmbedUnimplementedQueryServiceServer() {
}

func getStudents(querySql *QueryRequestInfo) []*Student {
	var students = make([]*Student, 2, 5)
	db, err := database.NewDBManager(utils.GetDatabaseURL())
	if err != nil {
		log.Fatalln("获取数据连接失败!!!!!")
	}

	rows, err := db.Query(querySql.GetInfo())
	if err != nil {
		log.Println("查询数据失败!!!!!")
	}

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		var age int
		var gender string
		var studentNumber string
		var grade string

		err := rows.Scan(&id, &name, &age, &gender, &studentNumber, &grade)
		if err != nil {
			log.Fatal(err.Error())
		}
		students = append(students, &Student{Id: fmt.Sprintf("%v", id), Name: name, Age: int32(age), Gender: gender, StudentNumber: studentNumber, Grade: grade})
	}

	return students
}
