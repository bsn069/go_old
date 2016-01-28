package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
)

const (
	CClientMsgHeaderSize uint = 1
)

// sUser sub class need imp
type iUserVirtual interface {
	getMsgBodyLen() uint16
	onReadOneMsg() error
	onDisconnect()
}

// bsn_net.IUser
type sUser struct {
	bsn_net.IUser
	iUserVirtual
	m_byMsgHeader []byte // need sub class init
	m_byMsgBody   []byte
}

func newUser(vIUser bsn_net.IUser, viUserVirtual iUserVirtual) (*sUser, error) {
	GLog.Debugln("newUser()")
	this := &sUser{}
	this.IUser = vIUser
	this.iUserVirtual = viUserVirtual
	return this, nil
}

func (this *sUser) run() {
	GLog.Debugln("run()")
	go this.runFunc()
}

func (this *sUser) runFunc() {
	GLog.Debugln("runFunc()")
	defer bsn_common.FuncGuard()
	defer this.Close()
	defer this.onDisconnect()

	vConn := this.GetConn()
	for {
		readLen, err := vConn.Read(this.m_byMsgHeader)
		if err != nil {
			GLog.Errorln(err)
			return
		}
		if readLen != len(this.m_byMsgHeader) {
			GLog.Errorln("not read all data")
			return
		}

		bodyLen := this.getMsgBodyLen()
		if bodyLen < bsn_msg.CMsgHeader_Size {
			GLog.Errorln("too short")
			return
		}
		if bodyLen > bsn_msg.CMsgSizeMax {
			GLog.Errorln("too long")
			return
		}

		this.m_byMsgBody = make([]byte, bodyLen)
		readLen, err = vConn.Read(this.m_byMsgBody)
		if err != nil {
			GLog.Errorln(err)
			return
		}
		if readLen != len(this.m_byMsgBody) {
			GLog.Errorln("not read all data")
			return
		}

		err = this.onReadOneMsg()
		if err != nil {
			GLog.Errorln(err)
			return
		}
	}
}
