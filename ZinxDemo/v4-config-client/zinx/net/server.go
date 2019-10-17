package net

import (
	"ZinxDemo/v4-config-client/zinx/config"
	"ZinxDemo/v4-config-client/zinx/iface"
	"fmt"
	"net"
	"strings"
)

//创建一个server 结构体
type Server struct {
	Name string
	TCPVersion string
	IP string
	Port uint32

	router iface.IRouter

}
func (s *Server)Serve(){
	fmt.Println("[Zinx Server Serve]")
	s.Start()
	//阻塞，不占用cpu，for循环会消耗cpu资源
	select {}
}
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
	var connid uint32
	go func() {
		for{
			tcpconn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Printf("tcplistener.AcceptTCP err:", err)
				return
			}
			fmt.Println("===> 新连接建立成功! <===")
			//3. 对conn进行处理（业务）：接收client信息，转换成大写返回
			conn :=NewConnection(tcpconn ,connid,server.router)
			connid++
			go conn.Start();
		}
	}()

	/*select {}*/
}
func (server *Server)Stop(){

}

//返回一个结构体对象
func NewServer(name string) iface.IServer{
	return &Server{
		IP: config.GlobalConfig.IP, //标识接收任意地址
		Port :config.GlobalConfig.Port,
		Name:       config.GlobalConfig.Name,
		TCPVersion: config.GlobalConfig.TcpVersion, //tcp , tcp4, tcp6
		router: &Router{},
	}
}

//定义一个函数
//func UserBussiness(conn iface.IConnection ,data []byte){
func UserBussiness(request iface.IRequest){
	data:=request.GetData()
	conn:=request.GetConn()
	//变成大写
	writeBackInfo := strings.ToUpper(string(data))
	//服务器发送给客户端
    conn.Send([]byte(writeBackInfo))
}

//注册到server结构体里一个方法
func (s *Server)AddRouter(router iface.IRouter){
	s.router =router
}



