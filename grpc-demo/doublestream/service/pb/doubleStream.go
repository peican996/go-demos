package pb

import (
	"io"
	"log"
)

var MyService = &myServiceServer{}

type myServiceServer struct{}

func (s *myServiceServer) MyMethod(stream MyService_MyMethodServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// 处理请求
		log.Printf("Received request: %s", req.Data)

		// 发送响应
		resp := &Response{Result: "Hello, " + req.Data}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func (s *myServiceServer) mustEmbedUnimplementedMyServiceServer() {

}
