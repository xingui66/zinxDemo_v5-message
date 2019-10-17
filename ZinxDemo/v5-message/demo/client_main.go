package main

import (
	net2 "ZinxDemo/v5-message/zinx/net"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	c,err := net.Dial("tcp","127.0.0.1:8848")
	if err!=nil {
		fmt.Println("net.Dail err :",err)
		return
	}

	data1 :=[]byte("helloworld")
	data2:= []byte ("jjjjjjjjj")
	data3:= []byte ("ooooooooooo")

	/*info :=append(data2,data1...)
	data:=append(info,data3...)*/
	//封包
	dp:=net2.NewDataPack()
	msg1,_:=dp.Pack(net2.NewMessage(uint32(len(data1)),1,data1))
	msg2,_:=dp.Pack(net2.NewMessage(uint32(len(data2)),1,data2))
	msg3,_:=dp.Pack(net2.NewMessage(uint32(len(data3)),1,data3))

	data := append(msg1,msg2...)
	data = append(data,msg3...)

	//多次向服务端传输数据，，并多次接收服务器传递过来的数据
	for {
		cnt,err := c.Write(data)
		if err!=nil {
			fmt.Println("c.Write err :",err)
			return
		}

		/*buffer := make([]byte,512)
		cnt,err = c.Read(buffer)*/
		headBuf :=make([]byte,8)
		cnt,err = io.ReadFull(c,headBuf)
		if err != nil {
		    fmt.Println("io.ReadFull err :",err)
		    return
		}
		fmt.Println("cnt:",cnt,"headBuf:",headBuf)
		dp:=net2.NewDataPack()
		message,err :=dp.UnPack(headBuf)
        if err != nil {
            fmt.Println("message err :",err)
            return
        }
		msgId :=message.GetMsgId()
		msgDataLen := message.GetDataLen()
		fmt.Println("dataLen:", msgDataLen, ", msgid:", msgId)
		if msgDataLen==0{
			fmt.Println("dataLen 长度为0，无需读取")
			return
		}

		dataBuf := make([]byte,msgDataLen)
		cnt,err = io.ReadFull(c,dataBuf)
		if err != nil {
		    fmt.Println("io.ReadFull dataBuf err:",err)
		    return
		}

		//fmt.Println("Client <=== Server, cnt:", cnt, ", data:", string(buffer[:cnt]))
		fmt.Println("cnt :", cnt, "dataBuf :", string(dataBuf)) //真实的数据
		//message.SetData(dataBuf)

		time.Sleep(1*time.Second)



	}

}


