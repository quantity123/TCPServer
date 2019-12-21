package workMgr

import "net"

type IWorkMgr interface {
	SetTcpConn(tcpConn *net.TCPConn)
}