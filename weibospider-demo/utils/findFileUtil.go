package utils

import (
	"demos/weibospider-demo/common"
	"fmt"
	"os"
	"path/filepath"
)

var path = ""

func FindFile(userId string) string {
	rootDir := common.Info.ScriptPath + "weibo\\" // 要查找的根目录
	err := filepath.Walk(rootDir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == userId+".txt" {
			path = file
			fmt.Println("目标文件路径:", file)
		}
		return nil
	})
	if err != nil {
		fmt.Println("遍历目录出错:", err)
	}
	return path
}
