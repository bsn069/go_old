package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"net"
)

type SSession struct {
	M_Conn net.Conn
}

func (this *SSession) Conn() net.Conn {
	return this.M_Conn
}

func (this *SSession) SetConn(vnetConn net.Conn) {
	this.M_Conn = vnetConn
}

func (this *SSession) Send(byData []byte) error {
	return Send(this.Conn(), byData)
}

func (this *SSession) SendString(strMsg string) error {
	return SendString(this.Conn(), strMsg)
}

func (this *SSession) SendMsgWithSMsgHeader(vTMsgType bsn_common.TMsgType, byMsg []byte) error {
	return SendMsgWithSMsgHeader(this.Conn(), vTMsgType, byMsg)
}

func (this *SSession) Recv(byData []byte) error {
	vReadLen := len(byData)

	vReadLen, err := this.Conn().Read(byData)
	if err != nil {
		if err.Error() == "EOF" {
			GSLog.Errorln("connect disconnect")
		} else {
			GSLog.Errorln("ReadMsg error: ", err)
		}
		return err
	}
	if vReadLen != len(byData) {
		return errors.New("not read all data")
	}
	return nil
}

func (this *SSession) RecvMsgWithSMsgHeader() (bsn_common.TMsgType, []byte, error) {
	byMsgHeader := make([]byte, bsn_msg.CSMsgHeader_Size)
	err := this.Recv(byMsgHeader)
	if err != nil {
		return 0, nil, err
	}
	vSMsgHeader := bsn_msg.NewMsgHeaderFromBytes(byMsgHeader)

	vu16Len := uint16(vSMsgHeader.Len())
	if vu16Len > 0 {
		byData := make([]byte, vu16Len)
		err = this.Recv(byData)
		if err != nil {
			return vSMsgHeader.Type(), nil, err
		}
	}

	return vSMsgHeader.Type(), nil, nil
}
