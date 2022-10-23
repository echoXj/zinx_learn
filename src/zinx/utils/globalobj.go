package utils

import (
	"encoding/json"
	"fmt"
	"go/src/zinx/ziface"
	"io/ioutil"
)

/*
	全局配置
*/
type GlobalObj struct {
	/*
		server相关
	*/
	TcpServer ziface.IServer // 当前zinx全局的server对象
	Host      string         //当前服务器监听IP
	TcpPort   int            //当前服务器监听端口
	Name      string         //当前服务器名称
	/*
		zinx相关
	*/
	Version        string // zinx版本号
	MaxConn        int    // 当前服务器主机最大链接数
	MaxPackageSize int32  // 当前zinx框架数据包最大值
}

/*
	定义一个全局的对外GlobalObj
*/
var GlobalObject *GlobalObj

/*
	从配置文件读取
*/
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("F:\\Golang\\go\\src\\myDemo\\zinxV0.4\\conf\\zinx.json")
	if err != nil {
		fmt.Println("Reload zinx.json err", err)
		panic(err)
	}
	// 将data中的数据解析到GlobalObj
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("Unmarshal zinx.json data err", err)
		panic(err)
	}
}

/*
	提供初始化GlobalObj的方法
*/
func init() {
	// 如果配置文件没有加载，就执行默认配置
	GlobalObject := &GlobalObj{
		Name:           "ZinxServerApp",
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Version:        "0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 尝试从配置文件读取
	GlobalObject.Reload()
}
