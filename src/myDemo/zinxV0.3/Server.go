package main

import (
	"fmt"
	"go/src/zinx/ziface"
	"go/src/zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (pr *PingRouter) PreHandle(request ziface.IReuest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil {
		fmt.Println("call back before error", err)
	}
}

// Test Handle
func (pr *PingRouter) Handle(request ziface.IReuest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back error", err)
	}
}

// Test PostHandle
func (pr *PingRouter) PostHandle(request ziface.IReuest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("call back after error", err)
	}
}

/*
	基于zinx开发一个服务端
*/
func main() {
	s := znet.NewServer("zinx_v0.3")
	// 给当前server增加一个router
	s.AddRouter(&PingRouter{})
	//启动Server
	s.Serve()
}
