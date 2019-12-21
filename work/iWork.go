package work

import "net"

type IWork interface {
	ContinueOprate(tcpConn *net.TCPConn)
}