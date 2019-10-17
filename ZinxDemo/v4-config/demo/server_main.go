package main

import (
	"ZinxDemo/v4-config/zinx/iface"
	"ZinxDemo/v4-config/zinx/net"
	"strings"
)

type TestRouter struct {
	net.Router
}

func (tr *TestRouter) Handle(request iface.IRequest){
	data:=request.GetData()
	conn:=request.GetConn()
	//变成大写
	writeBackInfo := strings.ToUpper(string(data))
	//服务器发送给客户端
	conn.Send([]byte(writeBackInfo))
}


func main() {
	server:=net.NewServer("zinx v1.0")
    //server.Start()
    server.AddRouter(&TestRouter{})
    server.Serve()
    //fmt.Println(server)
}