package iface

import "net"

type IConnection interface {
	GetTcpConn() *net.TCPConn
	GetConnId()uint32
	Start()
	Stop()
	Send([]byte,uint32)
}

//这里是函数
//type Callback func(IConnection, []byte)
type Callback func(IRequest)
