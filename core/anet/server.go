package anet

import (
	"ango/interface/inet"
	"ango/utils"
	"fmt"
	"net"
)

//iServer接口实现，定义一个Server服务类
type Server struct {
	Name        string                      //服务器的名称
	IPVersion   string                      //tcp4 or other
	IP          string                      //服务绑定的IP地址
	Port        int                         //服务绑定的端口
	msgHandler  inet.IMsgHandle             //当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	ConnMgr     inet.IConnManager           //当前Server的链接管理器
	OnConnStart func(conn inet.IConnection) //该Server的连接创建时Hook函数
	OnConnStop  func(conn inet.IConnection) //该Server的连接断开时的Hook函数
	packet      inet.Packet
}

/*
  创建一个服务器句柄
*/
func NewServer() inet.IServer {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()
	s := &Server{
		Name:       utils.GlobalObject.Name, //从全局参数获取
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,    //从全局参数获取
		Port:       utils.GlobalObject.TcpPort, //从全局参数获取
		msgHandler: NewMsgHandle(),             //msgHandler 初始化
		ConnMgr:    NewConnManager(),           //创建ConnManager
		packet:     NewDataPack(),              //封包拆包实例初始化
	}

	return s
}

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Ango Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//开启一个go去做服务端Listener业务
	go func() {
		//0 启动worker工作池机制
		s.msgHandler.StartWorkerPool()

		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		//已经监听成功
		fmt.Println("start Ango server ", s.Name, " success, now listening...")

		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				_ = conn.Close()
				continue
			}

			//3.3 处理该新连接请求的业务方法，此时应该有handler和conn是绑定的
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出，listener的go将会退出
	select {}
}

//AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgID uint32, router inet.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
}

//得到链接管理
func (s *Server) GetConnMgr() inet.IConnManager {
	return s.ConnMgr
}

//设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(inet.IConnection)) {
	s.OnConnStart = hookFunc
}

//设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(inet.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn inet.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn inet.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Server) Packet() inet.Packet {
	return s.packet
}
