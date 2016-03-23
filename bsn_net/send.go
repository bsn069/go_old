package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"net"
)

func Send(vnetConn net.Conn, byData []byte) error {
	vLen := len(byData)
	if vLen <= 0 {
		return nil
	}

	writeLen, err := vnetConn.Write(byData)
	if err != nil {
		return err
	}
	if writeLen != vLen {
		return errors.New("not write all data")
	}
	return nil
}

func SendString(vnetConn net.Conn, strMsg string) error {
	return Send(vnetConn, []byte(strMsg))
}

func SendMsgWithSMsgHeader(vnetConn net.Conn, vTMsgType bsn_common.TMsgType, byMsg []byte) error {
	vLen := len(byMsg)
	if vLen <= 0 {
		return nil
	}

	vMsg := bsn_msg.NewMsgHeader(vTMsgType, bsn_common.TMsgLen(vLen))

	byData := make([]byte, vLen+int(bsn_msg.CSMsgHeader_Size))
	vLen = int(vMsg.Serialize2Byte(byData))
	copy(byData[vLen:], byMsg)

	err := Send(vnetConn, byData)
	if err != nil {
		return err
	}

	return nil
}
