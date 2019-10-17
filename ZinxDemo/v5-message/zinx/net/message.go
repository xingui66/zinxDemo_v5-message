package net

import "ZinxDemo/v5-message/zinx/iface"

type Message struct {
	dataLen uint32
	msgId uint32
	data []byte
}

func NewMessage(len uint32,msgId uint32,data []byte) iface.IMessage{
	return &Message{
		dataLen:len,
		msgId:msgId,
		data:data,
	}
}

func (m *Message) GetDataLen() uint32 {
	return m.dataLen
}

func (m *Message) GetMsgId() uint32 {
	return m.msgId
}

func (m *Message) GetData() []byte {
	return m.data
}

func (m *Message) SetMsgId(msgId uint32) {
	m.msgId=msgId
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.dataLen=dataLen
}

func (m *Message)SetData(data []byte) {
	m.data=data
}









