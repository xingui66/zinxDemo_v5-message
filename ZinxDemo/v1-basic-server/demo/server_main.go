package main

import (
	"ZinxDemo/v1-basic-server/zinx/net"
)

func main() {
	server:=net.NewServer("zinx v1.0")
    //server.Start()
    server.Serve()
    //fmt.Println(server)
}