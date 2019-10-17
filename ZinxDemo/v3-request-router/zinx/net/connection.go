package net

import (
	"ZinxDemo/v3-request-router/zinx/iface"
	"fmt"
	"net"
)

type Connection struct {
	tcpConn *net.TCPConn
	connId uint32
	isClosed bool
	//callback iface.Callback
	route iface.IRouter


}
func NewConnection(conn *net.TCPConn,connid uint32,route iface.IRouter) iface.IConnection{
	return &Connection{
		tcpConn:conn,
		connId:connid,
		//callback:callback,
		route:route,
	}
}
func (c *Connection)GetTcpConn() *net.TCPConn{
	return c.tcpConn
}
func (c *Connection)GetConnId()uint32{
	return c.connId
}

///
func (c *Connection)Start() {
	for{
		buf :=make([]byte,512)
		cnt,err :=c.tcpConn.Read(buf)
		if err != nil {
			fmt.Printf("tcpconn.Read err:", err)
			return
		}
		fmt.Println("Server <=== Client, cnt:", cnt, "data:", string(buf[:cnt]))
		request:=NewRequest(c,buf[:cnt],uint32(cnt))
		//c.callback(request)
		c.route.Handle(request)
	}
}

//
//4. 关闭连接的方法：Stop
func (c *Connection) Stop() {
	fmt.Println("[Connection Stop called...]")
	//TODO
}

func (c *Connection)Send(data []byte){
	cnt,_ := c.GetTcpConn().Write([]byte(data))
	fmt.Println("Server ===> Client , cnt:", cnt)
}
