package server

import (
	"fmt"
	"net"
	"qtt/gameServer/common"
	"qtt/gameServer/workMgr"
	"sync"
)

type server struct {
	name string
	netWork string
	ip string
	port uint16
	workMgr workMgr.IWorkMgr
}

func (self *server)Launch(aWG *sync.WaitGroup) {
	defer aWG.Done()
	//写入执行日志到数据库里
	fmt.Println("qttGameServer Start Launching。")
	defer fmt.Println("qttGameServer Stop Launching。")
	addr := fmt.Sprintf("%s:%d", self.ip, self.port)
	tcpAddr, err := net.ResolveTCPAddr(self.netWork, addr)
	if err != nil {
		//写入错误日志到数据库里
		fmt.Println("net.ResolveTCPAddr err", err)
		return
	}
	tcpListener, err := net.ListenTCP(common.NetWork, tcpAddr)
	if err != nil {
		//写入错误日志到数据库里
		fmt.Println("net.ListenTCP err", err)
		return
	}
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("tcpListener.AcceptTCP err", err)
			continue
		}
		go self.workMgr.SetTcpConn(tcpConn)
	}
}

func NewServer() IServer {
	self := new(server)
	self.name = common.Name
	self.netWork = common.NetWork
	self.ip = common.IP
	self.port = common.Port
	self.workMgr = workMgr.NewWorkMgr()
	return self
}