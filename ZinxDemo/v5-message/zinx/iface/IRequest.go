package iface

type IRequest interface {
	GetConn() IConnection
	GetData() []byte
	GetLen() uint32

	GetMessage() IMessage
}
