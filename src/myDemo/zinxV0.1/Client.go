package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	// 直接连接远程服务器，的得到一个连接
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("net.dial err", err)
		return
	}

	// 连接调用write写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1..."))
		if err != nil {
			fmt.Println("conn.write err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.read buf err", err)
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
