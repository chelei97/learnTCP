package main

import (
	"TCPlearning/src/znet"
)

//Server 模块的测试函数
func main() {

	//1 创建一个server 句柄 s
	s := znet.NewServer("[zinx V0.2]")

	//2 开启服务
	s.Serve()
}