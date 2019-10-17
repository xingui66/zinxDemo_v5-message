package net

import (
	"ZinxDemo/v5-message/zinx/iface"
	"fmt"
	"io"
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
	for {
		headBuf := make([]byte ,8)
		cnt,err :=io.ReadFull(c.GetTcpConn(),headBuf)
		if err != nil {
			fmt.Println("io.ReadFull headbuf err :",err)
			return
		}
		fmt.Println("cnt :", cnt, "headBuf :", headBuf) //10/1
		dp :=NewDataPack()
		message,err:=dp.UnPack(headBuf)
		if err != nil {
			fmt.Println("",err)
			return
		}
		msgId := message.GetMsgId()
		dataLen:=message.GetDataLen()
		fmt.Println("dataLen:", dataLen, ", msgid:", msgId)
		if dataLen == 0 {
			fmt.Println("dataLen 长度为0，无需读取")
			return
		}

		//3. 根据解析出的数据长度len，读取真实的数据
		//&&&&&注意：这时候，读取的是//10/1后面的数据，因为前面read()了一部分了，再次read()就是后面的内容
		dataBuf := make([]byte,dataLen)
		cnt,err = io.ReadFull(c.GetTcpConn(),dataBuf);
		if err != nil {
			fmt.Println("io.ReadFull dataBuf err:", err)
			return
		}
		fmt.Println("cnt :", cnt, "dataBuf :", string(dataBuf)) //真实的数据
		message.SetData(dataBuf)

		request:=NewRequest(c,message)
		//c.callback(request)
		c.route.Handle(request)
	}

	/*for{
		buf :=make([]byte,512)
		cnt,err :=c.tcpConn.Read(buf)
		if err != nil {
			fmt.Println("tcpconn.Read err:", err)
			return
		}
		fmt.Println("Server <=== Client, cnt:", cnt, "data:", string(buf[:cnt]))
		request:=NewRequest(c,buf[:cnt],uint32(cnt))
		//c.callback(request)
		c.route.Handle(request)
	}*/
}

//
//4. 关闭连接的方法：Stop
func (c *Connection) Stop() {
	fmt.Println("[Connection Stop called...]")
	//TODO
}

func (c *Connection)Send(data[] byte,msgid uint32){
	dp :=NewDataPack()
	packInfo ,err := dp.Pack(NewMessage(uint32(len(data)),msgid,data))
	if err != nil {
	    fmt.Println("dp.Pack err :",err)
	    return
	}
	cnt,_ := c.tcpConn.Write([]byte(packInfo))
	fmt.Println("Server ====>Client,cnt:",cnt,".packinfo:",packInfo)
}


/*func (c *Connection)Send(data []byte){
	cnt,_ := c.GetTcpConn().Write([]byte(data))
	fmt.Println("Server ===> Client , cnt:", cnt)
}*/
