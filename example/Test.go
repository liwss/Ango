package main

import (
	"ango/core/anet"
	"fmt"
)

func main() {
	p := anet.NewEaiMsgPackage()
	p.SetMsgFlag(anet.MsgFlag)
	fmt.Println(p.GetMsgFlag())

	////创建一个存放bytes字节的缓冲
	//dataBuff := bytes.NewBuffer([]byte{})
	//
	////写dataLen
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.MsgFlag)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.Version)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.Finished)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.MsgType)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.MsgLength)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.AppId)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.TransId)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.TransNum)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.CodeTime)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.PositionId)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.UserId)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.StreamId)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.Priority)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.Transcode)
	//_ = binary.Write(dataBuff, binary.LittleEndian, p.MsgBody)
	//fmt.Println(dataBuff)
	//fmt.Println(dataBuff.Len())
	//
	//conn, err := net.Dial("tcp", "127.0.0.1:55555")
	//if err != nil {
	//	fmt.Println("client start err, exit!")
	//	return
	//}
	//_, err = conn.Write(dataBuff.Bytes())
	//if err != nil {
	//	fmt.Println("write error err ", err)
	//	return
	//}
	//
	////读出流中的head部分
	//headData := make([]byte, 80)
	//_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	//if err != nil {
	//	fmt.Println("read head error")
	//}
	//
	//fmt.Println(string(headData))
	//
	////读出流中的body部分
	//bodyData := make([]byte, 80)
	//_, err = io.ReadFull(conn, bodyData) //ReadFull 会把msg填充满为止
	//if err != nil {
	//	fmt.Println("read head error")
	//}
	//
	//fmt.Println(string(bodyData))
	//
	//
	//time.Sleep(100*time.Second)

}
