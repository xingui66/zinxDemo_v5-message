package main

import "ZinxDemo/v2-connection/zinx/net"

func main() {
	server:=net.NewServer("zinx v1.0")
    //server.Start()
    server.Serve()
    //fmt.Println(server)
}