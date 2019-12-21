package main

import (
	"fmt"
	"net"
	"qtt/gameServer/common"
	"qtt/gameServer/workMgr"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("net.ResolveTCPAddr err", err)
		return
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("net.DialTCP err", err)
		return
	}
	for {
		common.WriteUint16(tcpConn, uint16(workMgr.MSG_1))
		common.WriteUint16(tcpConn, uint16(6))
		v := common.ReadUint16(tcpConn)
		fmt.Println("add result:", v)

		common.WriteUint16(tcpConn, uint16(workMgr.MSG_2))
		var name string = "qtt"
		common.WriteString(tcpConn, name)
		s := common.ReadString(tcpConn)
		fmt.Println(s)

		time.Sleep(3 * time.Second)
	}
}
