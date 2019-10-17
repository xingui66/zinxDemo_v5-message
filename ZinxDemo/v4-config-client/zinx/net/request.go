package net

import "ZinxDemo/v4-config-client/zinx/iface"

//为了方便传递数据，我们将下面的数据封装成一个结构体
//func UserBussiness(conn iface.IConnection ,data []byte)

type Request struct {
	conn iface.IConnection
	data []byte
	len uint32
}

func NewRequest(conn iface.IConnection ,data []byte,len uint32) iface.IRequest{
  return &Request{
  	conn:conn,
  	len:len,
  	data:data,
  }
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