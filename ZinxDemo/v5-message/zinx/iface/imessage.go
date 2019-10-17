package iface

type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetMsgId(msgId uint32)
	SetDataLen(dataLen uint32)
	SetData(data []byte)
}
