package pb

import (
	"io"
	"log"
	"os"
)

var FileTransferServer = &fileTransferServer{}

type fileTransferServer struct{}

func (transferServer *fileTransferServer) TransferFile(stream FileTransferService_TransferFileServer) error {
	var fileName string
	var fileData []byte

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 文件传输完成，结束循环
			break
		}
		if err != nil {
			log.Fatalf("发生接收请求错误: %v", err)
		}
		// 根据请求类型进行处理
		if req.Filename != "" {
			// 接收到文件名
			fileName = req.Filename
		}
		log.Println(fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var num int
	fileData = make([]byte, 1024*1024*2)
	for {
		num, err = file.Read(fileData)
		log.Printf("=========: %v", num)
		if num == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// 发送文件数据给客户端
		err = stream.Send(&FileResponse{
			Filedata: fileData,
		})
		if err != nil {
			log.Fatalf("发送文件数据失败: %v", err)
		}
	}
	return nil
}

func (transferServer *fileTransferServer) mustEmbedUnimplementedFileTransferServiceServer() {

}
