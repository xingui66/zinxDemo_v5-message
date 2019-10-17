package net

import "ZinxDemo/v5-message/zinx/iface"

//为了方便传递数据，我们将下面的数据封装成一个结构体
//func UserBussiness(conn iface.IConnection ,data []byte)

type Request struct {
	conn iface.IConnection
	data []byte
	len uint32
	message iface.IMessage
}

//func NewRequest(conn iface.IConnection ,data []byte,len uint32) iface.IRequest{
func NewRequest(conn iface.IConnection ,message iface.IMessage) iface.IRequest{
  return &Request{
  	conn:conn,
  	message:message,
  }
}
func (r *Request) GetMessage() iface.IMessage {
	return r.message
}
func (r *Request)GetConn() iface.IConnection{
	return r.conn
}

func (r *Request)GetData()[]byte{
	return r.data
}

func (r *Request)GetLen()uint32{
	return r.len
}

