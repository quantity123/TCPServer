package workMgr

import (
	"net"
	"qtt/gameServer/common"
	"qtt/gameServer/work"
)

type MsgType uint16

const (
	MSG_1 MsgType = 1
	MSG_2 MsgType = 2
)

type workMgr struct {
	tcpConn *net.TCPConn
	workList map[MsgType]work.IWork
	workCount MsgType
}

func (self *workMgr)GetWork(aMsgID MsgType) work.IWork {
	if work, ok := self.workList[aMsgID]; ok {
		return work
	}
	return nil
}
/*
func (self *workMgr)ReadMsgIDD() MsgType {
	msgIdSize := make([]byte, 2)
	_, err := self.tcpConn.Read(msgIdSize)
	if err != nil {
		fmt.Println("tcpConn.Read msgIdSize err:", err)
		return 0
	}
	dataBuf := bytes.NewBuffer(msgIdSize)
	var msgID MsgType
	err = binary.Read(dataBuf, binary.LittleEndian, &msgID)
	if err != nil {
		fmt.Print("binary.Read msgID err:", err)
		return 0
	}
	return msgID

}
*/
func (self *workMgr)ReadMsgID() {
	for {
		msgID := MsgType(common.ReadUint16(self.tcpConn))
		switch {
		case msgID > 0:
			if work := self.GetWork(msgID); work != nil {
				work.ContinueOprate(self.tcpConn)
			}
		case msgID == 0:
		}
	}
}

func (self *workMgr)SetTcpConn(aTcpConn *net.TCPConn) {
	self.tcpConn = aTcpConn
	self.ReadMsgID()
}

func (self *workMgr)NewWorks() {
	var tmpWork work.IWork
	tmpWork = work.NewWork1()
	self.workList[MSG_1] = tmpWork
	self.workCount++
	tmpWork = work.NewWork2()
	self.workList[MSG_2] = tmpWork
	self.workCount++
}

func NewWorkMgr() IWorkMgr {
	self := new(workMgr)
	self.workCount = 0
	self.workList = make(map[MsgType]work.IWork)
	self.NewWorks()
	return self
}