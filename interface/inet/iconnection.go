/*
* @Author: liws
* @Date:   2021/11/20 14:53
* @Desc:
**/
package inet

import (
	"context"
	"net"
)

//定义连接接口
type IConnection interface {
	Start()                                      //启动连接，让当前连接开始工作
	Stop()                                       //停止连接，结束当前连接状态
	Context() context.Context                    //返回ctx，用于用户自定义的go程获取连接退出状态
	GetTCPConnection() *net.TCPConn              //从当前连接获取原始的socket TCPConn
	GetConnID() uint32                           //获取当前连接ID
	RemoteAddr() net.Addr                        //获取远程客户端地址信息
	StartWriter()                                //写消息Goroutine， 用户将数据发送给客户端
	StartReader()                                //读消息Goroutine
	SendMsg(msgId uint32, data []byte) error     //直接将Message数据发送数据给远程的TCP客户端
	SendBuffMsg(msgId uint32, data []byte) error //添加带缓冲发送消息接口
	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性
}

//定义一个统一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
