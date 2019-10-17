package net

import (
	"ZinxDemo/v5-message/zinx/iface"
	"bytes"
	"encoding/binary"
	"fmt"
)

type DataBack struct {}

func NewDataPack()*DataBack{
	return &DataBack{}
}

func (db *DataBack) GetHeadLen() uint32 {
	//msgId uint32（4个字节）  + DataLen uint32(4个字节)
	return 8
}

func (db *DataBack) Pack(message iface.IMessage) ([]byte, error) {
	dataBuff :=bytes.NewBuffer([]byte{})
	err :=binary.Write(dataBuff,binary.LittleEndian,message.GetDataLen())
	if err != nil {
		return nil,err
	}

	binary.Write(dataBuff,binary.LittleEndian,message.GetMsgId())
	if err != nil {
	    return nil ,err
	}
	binary.Write(dataBuff,binary.LittleEndian,message.GetData());
	if err != nil {
		return nil ,err
	}
	return dataBuff.Bytes(),err
}

func (db *DataBack) UnPack(data []byte) (iface.IMessage, error) {
	var message Message
	reader := bytes.NewBuffer(data)

	err :=binary.Read(reader,binary.LittleEndian,&message.dataLen)
	if err != nil {
	    fmt.Println("&message.dataLen:",err)
	    return  nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &message.msgId)
	if err != nil {
		fmt.Println("binary.Read message msgid err:", err)
		return nil, err
	}

	//Message len: 10, msgid : 1, data :[]byte{}

	return &message, nil
}



