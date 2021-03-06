/*
* @Author: liws
* @Date:   2021/11/20 14:53
* @Desc:
**/
package inet

/*
	消息管理抽象层
*/
type IMsgHandle interface {
	DoMsgHandler(request IRequest)                        //马上以非阻塞方式处理消息
	AddRouter(msgId uint32, router IRouter)               //为消息添加具体的处理逻辑
	StartOneWorker(workerID int, taskQueue chan IRequest) //启动一个Worker工作流程
	StartWorkerPool()                                     //启动worker工作池
	SendMsgToTaskQueue(request IRequest)                  //将消息交给TaskQueue,由worker进行处理
}
