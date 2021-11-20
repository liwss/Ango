/*
* @Author: liws
* @Date:   2021/11/20 14:53
* @Desc:   将请求的一个消息封装到message中，定义抽象层接口
**/
package inet

type IEaiMessage interface {
	GetMsgFlag() string          //获取消息起始标志
	GetVersion() string          //获取消息版本号
	GetFinished() byte           //获取是否最后一包
	GetMsgType() byte            //获取消息类型
	GetMsgLength() (int, error)  //获取消息体长度
	GetAppId() string            //获取事务类型
	GetTransId() (int, error)    //获取事务标志
	GetTransNum() (int, error)   //获取事务分组序号
	GetCodeTime() (int, error)   //获取生存期
	GetPositionId() (int, error) //获取任务请求客户标识
	GetUserId() string           //获取用户ID
	GetStreamId() (int, error)   //获取通讯描述符序号
	GetPriority() (int, error)   //获取优先级
	GetTransCode() string        //获取交易代码
	GetMsgBody() []byte          //获取消息体

	SetMsgFlag(string)   //设置消息起始标志
	SetVersion(string)   //设置消息版本号
	SetFinished(byte)    //设置是否最后一包
	SetMsgType(byte)     //设置消息类型
	SetMsgLength(int)    //设置消息体长度
	SetAppId(string)     //设置事务类型
	SetTransId(int)      //设置事务标志
	SetTransNum(int)     //设置事务分组序号
	SetCodeTime(int)     //设置生存期
	SetPositionId(int)   //设置任务请求客户标识
	SetUserId(string)    //设置用户ID
	SetStreamId(int)     //设置通讯描述符序号
	SetPriority(int)     //设置优先级
	SetTransCode(string) //设置交易代码
	SetMsgBody([]byte)   //设置消息体
}
