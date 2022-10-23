package ziface

/*
	封装请求的链接和数据
*/
type IReuest interface {
	// 获取链接
	GetConnection() IConnection

	// 获取数据
	GetData() []byte
}
