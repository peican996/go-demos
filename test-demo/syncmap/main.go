package main

import (
	"fmt"
	"time"
)

func main() {
	demo := make(map[int]string)

	go func() {
		for i := 0; i < 1000; i++ {
			demo[i] = "a"
		}
	}()
	time.Sleep(time.Duration(10) * time.Second)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(demo[i])
		}
	}()

}
