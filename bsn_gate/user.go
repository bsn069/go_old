package bsn_gate

import (
	"errors"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	"net"
)

type TUserId uint16
type IUser interface {
	GetId() TUserId
	GetConn() net.Conn
	ReadMsgHeader() error
	ReadMsgBody() error
}

// IUser
type SUser struct {
	M_TUserId     TUserId
	M_Conn        net.Conn
	M_byMsgHeader []byte
	M_byMsgBody   []byte
	M_bClose      bool
}

func newUser() (*SUser, error) {
	GLog.Debugln("newUser()")
	this := &SUser{}
	return this, nil
}

func (this *SUser) GetId() TUserId {
	return this.M_TUserId
}

func (this *SUser) SetId(vTUserId TUserId) error {
	this.M_TUserId = vTUserId
	return nil
}

func (this *SUser) GetConn() net.Conn {
	return this.M_Conn
}

func (this *SUser) SetConn(vConn net.Conn) error {
	this.M_Conn = vConn
	return nil
}

func (this *SUser) ReadMsgHeader() error {
	readLen, err := this.M_Conn.Read(this.M_byMsgHeader)
	if err != nil {
		return err
	}
	if readLen != len(this.M_byMsgHeader) {
		return errors.New("not read all data")
	}
	return nil
}

func (this *SUser) ReadMsgBody() error {
	readLen, err := this.M_Conn.Read(this.M_byMsgBody)
	if err != nil {
		return err
	}
	if readLen != len(this.M_byMsgBody) {
		return errors.New("not read all data")
	}
	return nil
}
