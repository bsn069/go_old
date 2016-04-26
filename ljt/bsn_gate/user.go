package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	"net"
)

// IGateUser
type SUser struct {
	M_TUserId     bsn_common.TGateUserId
	M_Conn        net.Conn
	M_byMsgHeader []byte
	M_byMsgBody   []byte
	M_bClose      bool
}

func NewUser() (*SUser, error) {
	GSLog.Debugln("NewUser()")
	this := &SUser{}
	return this, nil
}

func (this *SUser) Id() bsn_common.TGateUserId {
	return this.M_TUserId
}

func (this *SUser) SetId(vTUserId bsn_common.TGateUserId) error {
	this.M_TUserId = vTUserId
	return nil
}

func (this *SUser) Conn() net.Conn {
	return this.M_Conn
}

func (this *SUser) SetConn(vConn net.Conn) error {
	this.M_Conn = vConn
	return nil
}

func (this *SUser) ReadMsgHeader() error {
	readLen, err := this.Conn().Read(this.M_byMsgHeader)
	if err != nil {
		return err
	}
	if readLen != len(this.M_byMsgHeader) {
		return errors.New("not read all data")
	}
	return nil
}

func (this *SUser) ReadMsgBody() error {
	readLen, err := this.Conn().Read(this.M_byMsgBody)
	if err != nil {
		return err
	}
	if readLen != len(this.M_byMsgBody) {
		return errors.New("not read all data")
	}
	return nil
}

func (this *SUser) Send(byData []byte) error {
	this.Conn().Write(byData)
	return nil
}

func (this *SUser) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.Conn().Close()
	return nil
}
