package net

import (
	"ZinxDemo/v2-connection/zinx/iface"
	"fmt"
	"net"
)

type Connection struct {
	tcpConn *net.TCPConn
	connId uint32
	isClosed bool
	callback iface.Callback
}
func NewConnection(conn *net.TCPConn,connid uint32,callback iface.Callback) iface.IConnection{
	return &Connection{
		tcpConn:conn,
		connId:connid,
		callback:callback,
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
		c.callback(c,buf[:cnt])
	}
}

//
//4. 关闭连接的方法：Stop
func (c *Connection) Stop() {
	fmt.Println("[Connection Stop called...]")
	//TODO
}
