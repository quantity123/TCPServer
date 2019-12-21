package work

import (
	"net"
	"qtt/gameServer/common"
)

type work2 struct {
	anser string
}

func (self *work2)ContinueOprate(tcpConn *net.TCPConn) {
	name := common.ReadString(tcpConn)
	s := self.anser + " " + name
	common.WriteString(tcpConn, s)
}

func NewWork2() IWork {
	self := new(work2)
	self.anser = "hello"
	return self
}