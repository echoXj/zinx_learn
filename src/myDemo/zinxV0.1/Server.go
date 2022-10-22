package main

import "go/src/zinx/znet"

/*
	基于zinx开发一个服务端
*/
func main() {
	s := znet.NewServer("zinx_v0.1")
	s.Serve()

}
