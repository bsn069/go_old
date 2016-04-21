package bsn_net

import (
	"bsn_define"
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
	"net"
)

type SSession struct {
	M_Conn net.Conn
}

func NewSSession() (*SSession, error) {
	GSLog.Debugln("NewSSession")

	this := &SSession{}

	return this, nil
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

func (this *SSession) Ping(byMsg []byte) error {
	return this.SendMsgWithSMsgHeader(bsn_common.TMsgType(bsn_define.ECmd_Cmd_Ping), byMsg)
}

func (this *SSession) Pong(byMsg []byte) error {
	return this.SendMsgWithSMsgHeader(bsn_common.TMsgType(bsn_define.ECmd_Cmd_Pong), byMsg)
}

func (this *SSession) SendPbMsgWithSMsgHeader(vTMsgType bsn_common.TMsgType, iMessage proto.Message) error {
	b, err := proto.Marshal(iMessage)
	if err != nil {
		GSLog.Errorln(err)
		return err
	}
	return this.SendMsgWithSMsgHeader(vTMsgType, b)
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
