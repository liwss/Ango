package anet

const (
	MsgFlag     = "`DC`" // 消息起始标志
	Version     = "01"   // 版本号
	Finished    = '1'    // 1:最后一包
	NotFinished = '0'    // 0:非最后一包
	MsgTypeReq  = '3'    // 请求包标志
	MsgTypeRsp  = '4'    // 响应包标志
)

const (
	MsgFlagLen    = 4  // 消息起始标志长度
	VersionLen    = 2  // 版本号长度
	MsgLen        = 4  // 消息体长度，最大8000
	AppIdLen      = 6  // 事务类型长度
	TransIdLen    = 10 // 事务标志长度
	transNumLen   = 6  // 事务内分组序号长度
	CodeTimeLen   = 6  // 生存期长度
	PositionIdLen = 8  // 任务请求客户标识长度
	UserIdLen     = 8  // 用户ID长度
	StreamIdLen   = 8  // 通讯描述符序号长度
	PriorityLen   = 2  // 优先级长度
	TranscodeLen  = 14 // 交易代码长度
)
