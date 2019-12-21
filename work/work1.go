package work

import (
	"net"
	"qtt/gameServer/common"
)

type work1 struct {
	A uint16
}

func (self *work1)addition(B uint16) uint16 {
	return self.A + B
}

func (self *work1)ContinueOprate(tcpConn *net.TCPConn) {
	B := common.ReadUint16(tcpConn)
	addRst := self.addition(B)
	common.WriteUint16(tcpConn, addRst)
}

func NewWork1() IWork {
	self := new(work1)
	self.A = 10
	return self
}