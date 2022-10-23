package znet

import (
	"go/src/zinx/ziface"
)

type Request struct {
	// 请求的链接(已经和客户端建立好的链接)
	conn ziface.IConnection

	// 请求的参数
	data []byte
}

// 获取当前请求的来链接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取当前请求的数据
func (r *Request) GetData() []byte {
	return r.data
}
