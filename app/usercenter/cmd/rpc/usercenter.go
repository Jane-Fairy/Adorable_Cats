package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var configFile = flag.String("f", "app/usercenter/cmd/rpc/etc/usercenter.yaml", "the config file")

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 02132
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine

	fmt.Print("all done")
}

var lg chan string

func BroadcastOne(roleId string) {
	timeout := time.NewTimer(10 * time.Second)
	for {
		select {
		case lg <- roleId:
			log.Print("Received: ", <-lg)
			time.Sleep(2 * time.Second)
		case <-timeout.C:
			log.Print("Timeout")
			// 重置定时器

			timeout.Reset(10 * time.Second)
		}
	}
}
