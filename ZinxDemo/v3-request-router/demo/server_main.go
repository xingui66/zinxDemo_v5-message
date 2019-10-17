package main

import (
	"ZinxDemo/v3-request-router/zinx/iface"
	"ZinxDemo/v3-request-router/zinx/net"
)

type TestRouter struct {
	net.Router
}

func (tr *TestRouter) Handle(request iface.IRequest){

}


func main() {
	server:=net.NewServer("zinx v1.0")
    //server.Start()
    server.Serve()
    //fmt.Println(server)
}