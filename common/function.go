package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func ReadUint16(aTcpConn *net.TCPConn) uint16 {
	buf := make([]byte, 2)
	_, err := aTcpConn.Read(buf)
	if err != nil {
		fmt.Println("tcpConn.Read buf err:", err)
		return 0
	}
	dataBuf := bytes.NewBuffer(buf)
	var v uint16
	err = binary.Read(dataBuf, binary.LittleEndian, &v)
	if err != nil {
		fmt.Print("binary.Read dataBuf err:", err)
		return 0
	}
	return v
}

func WriteUint16(aTcpConn *net.TCPConn, aV uint16) {
	var buf []byte
	dataBuf := bytes.NewBuffer(buf)
	err := binary.Write(dataBuf, binary.LittleEndian, aV)
	if err != nil {
		fmt.Print("binary.Write dataBuf err:", err)
		return
	}
	_, err = aTcpConn.Write(dataBuf.Bytes())
	if err != nil {
		fmt.Print("aTcpConn.Write dataBuf err:", err)
	}
}

func ReadString(aTcpConn *net.TCPConn) string {
	strLen := ReadUint16(aTcpConn)
	if strLen <= 0 {
		return ""
	}
	strBuf := make([]byte, strLen)
	_, err := aTcpConn.Read(strBuf)
	if err != nil {
		fmt.Print("ReadString s err:", err)
		return ""
	}
	dataBuf := bytes.NewBuffer(strBuf)
	return dataBuf.String()
}

func WriteString(aTcpConn *net.TCPConn, aS string) {
	dataBuf := bytes.NewBufferString(aS)
	WriteUint16(aTcpConn, uint16(dataBuf.Len()))
	_, err := aTcpConn.Write(dataBuf.Bytes())
	if err != nil {
		fmt.Print("WriteString dataBuf err:", err)
	}
}