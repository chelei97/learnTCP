package znet

import (
	"learnZinx/src/ziface"
	"fmt"
	"net"
)

//iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的ip
	IP string
	//服务器监听的端口
	port int
}

func (s *Server) Start(){
	fmt.Printf("[Start] Server Listener at IP : %s, Port %d, is starting\n", s.IP, s.port)
	go func() {
		//获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}
		//监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err ", err)
		}
		fmt.Println("start Zinx server, ", s.Name, " success, Listening...")
		//阻塞等待客户端连接，处理客户端连接业务
		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			//已经与客户端建立连接，做一些业务，做一个最基本的最大512字节长度的回显业务
			go func(){
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}

					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					//回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop(){

}

func (s *Server) Serve(){
	//启动server的服务功能
	s.Start()

	//做一些启动服务器之后的额外业务

	//阻塞状态
	select {

	}
}

//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name : name,
		IPVersion : "tcp4",
		IP : "0.0.0.0",
		port : 8999,
	}
	return s
}