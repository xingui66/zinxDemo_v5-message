package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestNewDataPack(t *testing.T) {

	go func() {
		l,err := net.Listen("tcp4",":8848")
		if err != nil {
		    fmt.Println("net.Listen err:",err)
		    return
		}

		conn,err := l.Accept()
		if err != nil {
		    fmt.Println("l.Accept err :",err)
		    return
		}

		for {
			headBuf := make([]byte ,8)
			cnt,err :=io.ReadFull(conn,headBuf)
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
			cnt,err = io.ReadFull(conn,dataBuf);
			if err != nil {
				fmt.Println("io.ReadFull dataBuf err:", err)
				return
			}
			fmt.Println("cnt :", cnt, "dataBuf :", string(dataBuf)) //真实的数据
		}

	}()

	go func() {
		c,err := net.Dial("tcp4",":8048")
		if err != nil {
		    fmt.Println("netDial err ;",err)
		    return
		}

		data1 := []byte("helloworld")
		data2 := []byte("喵喵喵")
		data3 := []byte("大傻逼")

		dp:=NewDataPack()
		msg1,_:= dp.Pack(NewMessage(uint32(len(data1)),1,data1))
		msg2,_:= dp.Pack(NewMessage(uint32(len(data2)),2,data2))
		msg3,_:= dp.Pack(NewMessage(uint32(len(data3)),3,data3))

		data:=append(msg1,msg2...)
		data=append(msg3,msg2...)
        cnt,err:=c.Write(data)
        if err != nil {
            fmt.Println("c.Write err :",err)
            return
        }

		fmt.Println("Client ===> Server, cnt:", cnt, ", data:", string(data[:cnt]), "binary data:", data[:cnt])

		buffer:=make([]byte,512)
        cnt,err = c.Read(buffer)
        if err != nil {
            fmt.Println("c.Read err:",err)
            return
        }

		fmt.Println("Client <=== Server, cnt:", cnt, ", data:", string(buffer[:cnt]))
		
	}()

	select {
	}
}

/*func TestNewDataPack(t *testing.T) {
	fmt.Println("heheheh")
}*/
