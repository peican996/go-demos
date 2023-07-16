package service

import (
	"fmt"
	"os/exec"
)

func WeiboService(configFile string) {
	cmd := exec.Command("python", "-u", "-X", "utf8", "-m", "weibo_spider", "--config_path="+configFile)
	cmd.Dir = "D:\\code\\python\\weiboSpider\\"
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdoutPipe.Read(buf)
			if err != nil {
				fmt.Println(err)
			}
			output := string(buf[:n])
			if output == "" {
				break
			}
			fmt.Println(output)
			//if strings.Contains(output, "完毕") {
			//	fmt.Println(output)
			//}
		}
	}()
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
