package bsn_common

import (
	"net"
)

func WriteMsg(conn net.Conn, u16Type uint16, byData []byte) (err error) {
	var u16Len uint16 = uint16(len(byData))
	msgHeader := newMsgHeader(u16Type, u16Len)
	_, err = conn.Write(msgHeader.bytes())
	if err != nil {
		return
	}

	if u16Len > 0 {
		_, err = conn.Write(byData)
		if err != nil {
			return
		}
	}

	return nil
}

func ReadMsg(conn net.Conn) (err error, u16Type uint16, byData []byte) {
	byMsgHeader := make([]byte, cMsgHeader_Size)
	_, err = conn.Read(byMsgHeader)
	if err != nil {
		return
	}
	msgHeader := newMsgHeaderFromBytes(byMsgHeader)

	var u16Len uint16 = msgHeader.getLen()
	var byMsgBody []byte
	if u16Len > 0 {
		byMsgBody = make([]byte, u16Len)
		_, err = conn.Read(byMsgBody)
		if err != nil {
			return
		}
	}

	return nil, msgHeader.getType(), byMsgBody
}
