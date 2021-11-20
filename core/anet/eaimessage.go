/*
* @Author: liws
* @Date:   2021/11/20 14:53
* @Desc:
**/
package anet

import (
	"fmt"
	"strconv"
)

type EaiMessage struct {
	msgFlag    [MsgFlagLen]byte    //消息的起始标志 "`DC`"
	version    [VersionLen]byte    //版本号 "01"
	finished   byte                //是否是最后一包  1:是 0:否
	msgType    byte                //消息类型 3:请求消息 4:响应消息
	msgLength  [MsgLen]byte        //消息长度，不足4位左补零，最大8000字节，超过8k多包传输
	appId      [AppIdLen]byte      //事务类型
	transId    [TransIdLen]byte    //事务标志，唯一的序号，不足10位左补零占位长度
	transNum   [transNumLen]byte   //事务内分组序号（必须连续）,不足6位左补零占位长度
	codeTime   [CodeTimeLen]byte   //生存期,单位秒(或应答标识),不足6位左补零占位长度
	positionId [PositionIdLen]byte //任务请求客户标识
	userId     [UserIdLen]byte     //通信代理客户标识
	streamId   [StreamIdLen]byte   //通信描述符序号
	priority   [PriorityLen]byte   //优先级
	transcode  [TranscodeLen]byte  //交易代码,不足14位左空格补位
	msgBody    []byte              //消息内容
}

//创建一个EaiMessage消息包
func NewEaiMsgPackage() *EaiMessage {
	return &EaiMessage{}
}

//获取消息起始标志
func (msg *EaiMessage) GetMsgFlag() string {
	return string(msg.msgFlag[:])
}

//设置消息起始标志
func (msg *EaiMessage) SetMsgFlag(msgFlag string) {
	temp := []byte(msgFlag)
	length := len(temp)
	if length == MsgFlagLen {
		copy(msg.msgFlag[:], temp[:length])
	} else if length < MsgFlagLen {
		copy(msg.msgFlag[MsgFlagLen-length:], temp[:length])
	} else {
		fmt.Println("msgFlag is too bigger")
	}
}

//获取版本号
func (msg *EaiMessage) GetVersion() string {
	return string(msg.version[:])
}

//设置版本号
func (msg *EaiMessage) SetVersion(version string) {
	temp := []byte(version)
	length := len(temp)
	if length == VersionLen {
		copy(msg.version[:], temp[:length])
	} else if length < VersionLen {
		copy(msg.version[VersionLen-length:], temp[:length])
	} else {
		fmt.Println("version is too bigger")
	}
}

//获取是否最后包
func (msg *EaiMessage) GetFinished() byte {
	return msg.finished
}

//设置是否最后包
func (msg *EaiMessage) SetFinished(finished byte) {
	msg.finished = finished
}

//获取消息类型
func (msg *EaiMessage) GetMsgType() byte {
	return msg.msgType
}

//设置消息类型
func (msg *EaiMessage) SetMsgType(msgType byte) {
	msg.msgType = msgType
}

//获取消息体长度
func (msg *EaiMessage) GetMsgLength() (int, error) {
	return strconv.Atoi(string(msg.msgLength[:]))
}

//设置消息体长度
func (msg *EaiMessage) SetMsgLength(msgLength int) {
	temp := []byte(strconv.Itoa(msgLength))
	length := len(temp)
	if length == MsgLen {
		copy(msg.msgLength[:], temp[:length])
	} else if length < MsgLen {
		copy(msg.msgLength[MsgLen-length:], temp[:length])
	} else {
		fmt.Println("msgLength is too bigger")
	}
}

// 获取事务类型
func (msg *EaiMessage) GetAppId() string {
	return string(msg.appId[:])
}

// 设置事务类型
func (msg *EaiMessage) SetAppId(appId string) {
	temp := []byte(appId)
	length := len(temp)
	if length == AppIdLen {
		copy(msg.appId[:], temp[:length])
	} else if length < AppIdLen {
		copy(msg.appId[AppIdLen-length:], temp[:length])
	} else {
		fmt.Println("appIdLen is too bigger")
	}
}

//获取事务标志
func (msg *EaiMessage) GetTransId() (int, error) {
	return strconv.Atoi(string(msg.transId[:]))
}

//设置事务标志
func (msg *EaiMessage) SetTransId(tansId int) {
	temp := []byte(strconv.Itoa(tansId))
	length := len(temp)
	if length == TransIdLen {
		copy(msg.transId[:], temp[:length])
	} else if length < TransIdLen {
		copy(msg.transId[TransIdLen-length:], temp[:length])
	} else {
		fmt.Println("transIdLen is too bigger")
	}
}

//获取事务分组序号
func (msg *EaiMessage) GetTransNum() (int, error) {
	return strconv.Atoi(string(msg.transNum[:]))
}

//设置事务分组序号
func (msg *EaiMessage) SetTransNum(transNum int) {
	temp := []byte(strconv.Itoa(transNum))
	length := len(temp)
	if length == transNumLen {
		copy(msg.transNum[:], temp[:length])
	} else if length < transNumLen {
		copy(msg.transNum[transNumLen-length:], temp[:length])
	} else {
		fmt.Println("transNumLen is too bigger")
	}
}

//获取生存期
func (msg *EaiMessage) GetCodeTime() (int, error) {
	return strconv.Atoi(string(msg.codeTime[:]))
}

//设置生存期
func (msg *EaiMessage) SetCodeTime(codeTime int) {
	temp := []byte(strconv.Itoa(codeTime))
	length := len(temp)
	if length == CodeTimeLen {
		copy(msg.codeTime[:], temp[:length])
	} else if length < CodeTimeLen {
		copy(msg.codeTime[CodeTimeLen-length:], temp[:length])
	} else {
		fmt.Println("codeTimeLen is too bigger")
	}
}

// 获取任务请求客户标识
func (msg *EaiMessage) GetPositionId() (int, error) {
	return strconv.Atoi(string(msg.positionId[:]))
}

// 设置任务请求客户标识
func (msg *EaiMessage) SetPositionId(positionId int) {
	temp := []byte(strconv.Itoa(positionId))
	length := len(temp)
	if length == PositionIdLen {
		copy(msg.positionId[:], temp[:length])
	} else if length < PositionIdLen {
		copy(msg.positionId[PositionIdLen-length:], temp[:length])
	} else {
		fmt.Println("positionIdLen is too bigger")
	}
}

//获取用户ID
func (msg *EaiMessage) GetUserId() string {
	return string(msg.userId[:])
}

//设置用户ID
func (msg *EaiMessage) SetUserId(userId string) {
	temp := []byte(userId)
	length := len(temp)
	if length == UserIdLen {
		copy(msg.userId[:], temp[:length])
	} else if length < UserIdLen {
		copy(msg.userId[UserIdLen-length:], temp[:length])
	} else {
		fmt.Println("userIdLen is too bigger")
	}
}

//获取通讯描述符序号
func (msg *EaiMessage) GetStreamId() (int, error) {
	return strconv.Atoi(string(msg.streamId[:]))
}

//设置通讯描述符序号
func (msg *EaiMessage) SetStreamId(streamId int) {
	temp := []byte(strconv.Itoa(streamId))
	length := len(temp)
	if length == StreamIdLen {
		copy(msg.streamId[:], temp[:length])
	} else if length < StreamIdLen {
		copy(msg.streamId[StreamIdLen-length:], temp[:length])
	} else {
		fmt.Println("streamIdLen is too bigger")
	}
}

//获取优先级
func (msg *EaiMessage) GetPriority() (int, error) {
	return strconv.Atoi(string(msg.priority[:]))
}

//设置优先级
func (msg *EaiMessage) SetPriority(priority int) {
	temp := []byte(strconv.Itoa(priority))
	length := len(temp)
	if length == PriorityLen {
		copy(msg.priority[:], temp[:length])
	} else if length < PriorityLen {
		copy(msg.priority[PriorityLen-length:], temp[:length])
	} else {
		fmt.Println("priorityLen is too bigger")
	}
}

//获取交易代码
func (msg *EaiMessage) GetTransCode() string {
	return string(msg.transcode[:])
}

//设置交易代码
func (msg *EaiMessage) SetTransCode(transcode string) {
	temp := []byte(transcode)
	length := len(temp)
	if length == TranscodeLen {
		copy(msg.transcode[:], temp[:length])
	} else if length < TranscodeLen {
		copy(msg.transcode[TranscodeLen-length:], temp[:length])
	} else {
		fmt.Println("transcodeLen is too bigger")
	}
}

//获取消息体
func (msg *EaiMessage) GetMsgBody() []byte {
	return msg.msgBody
}

//设置消息体
func (msg *EaiMessage) SetMsgBody(message []byte) {
	msg.msgBody = message
}
