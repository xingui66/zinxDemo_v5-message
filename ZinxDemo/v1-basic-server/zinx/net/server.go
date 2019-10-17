package net

import (
	"ZinxDemo/v1-basic-server/zinx/iface"
	"fmt"
	"net"
	"strings"
)

//创建一个server 结构体
type Server struct {
	Name string
	TCPVersion string
	IP string
	Port int
}
//绑定两个方法
func (server *Server)Start(){
	fmt.Println("[Zinx Server Start....]")
	//TODO , 标识以后在实现具体的逻辑
	//旧版本：
	//l:=net.Listen("tcp", ":9999")
	//c := l.Accept()
	//c.Read()
	addr:=fmt.Sprintf("%s:%d",server.IP,server.Port)

	tcpAddr ,err := net.ResolveTCPAddr(server.TCPVersion,addr)
	if err != nil {
		fmt.Printf("net.ResolveTCPAddr err:", err)
		return
	}
    tcpListener,err := net.ListenTCP(server.TCPVersion,tcpAddr)
	if err != nil {
		fmt.Printf("net.ListenTCP err:", err)
		return
	}

	//2. 建立连接Accept
	//没有办法接收多个链接，后续使用goroutine处理即可
	go func() {
		for{
			tcpconn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Printf("tcplistener.AcceptTCP err:", err)
				return
			}
			fmt.Println("===> 新连接建立成功! <===")
			//3. 对conn进行处理（业务）：接收client信息，转换成大写返回
			go func() {
				for{
					buf :=make([]byte,512)
					cnt,err :=tcpconn.Read(buf)
					if err != nil {
						fmt.Printf("tcpconn.Read err:", err)
						return
					}
					fmt.Println("Server <=== Client, cnt:", cnt, "data:", string(buf[:cnt]))
					//变成大写
					writeBackInfo := strings.ToUpper(string(buf[:cnt]))
					//服务器发送给客户端
					cnt, _ = tcpconn.Write([]byte(writeBackInfo))
					fmt.Println("Server ===> Client , cnt:", cnt)
				}
			}()
		}
	}()

	/*select {}*/
}
func (server *Server)Stop(){

}
func (s *Server)Serve(){
	fmt.Println("[Zinx Server Serve]")
	s.Start()
	//阻塞，不占用cpu，for循环会消耗cpu资源
	select {}
}

//返回一个结构体对象
func NewServer(name string) iface.IServer{
	return &Server{
		IP:         "0.0.0.0", //标识接收任意地址
		Port:       8848,
		Name:       name,
		TCPVersion: "tcp4", //tcp , tcp4, tcp6
	}
}





