package znet

import "learnTCP/interface"

type Request struct {
	conn _interface.IConnection //已经和客户端建立好的 链接
	msg  _interface.IMessage    //客户端请求的数据
}
//获取请求连接信息
func(r *Request) GetConnection() _interface.IConnection {
	return r.conn
}
//获取请求消息的数据
func(r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}