package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

// IUser
type SClientUser struct {
	*SUser
	M_SClientUserMgr *SClientUserMgr
}

func newClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("newClientUser")
	this := &SClientUser{}

	vSUser, err := newUser()
	if err != nil {
		return nil, err
	}
	this.SUser = vSUser
	this.M_SClientUserMgr = vSClientUserMgr

	this.M_byMsgHeader = make([]byte, CSClientMsg_Size)

	return this, nil
}

func (this *SClientUser) Run() {
	go this.runImp()
}

func (this *SClientUser) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.M_Conn.Close()
	return nil
}

func (this *SClientUser) Send(byData []byte) error {
	this.M_Conn.Write(byData)
	return nil
}

func (this *SClientUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		this.M_SClientUserMgr.DelUser(this)
	}()

	var vClientMsg SClientMsg
	for !this.M_bClose {
		err := this.ReadMsgHeader()
		if err != nil {
			GSLog.Errorln(err)
			return
		}
		vClientMsg.DeSerialize(this.M_byMsgHeader)

		vLen := vClientMsg.Len()
		if vLen < bsn_msg.CSMsgHeader_Size {
			GSLog.Errorln("too short")
			return
		}

		this.M_byMsgBody = make([]byte, vLen)
		err = this.ReadMsgBody()
		if err != nil {
			GSLog.Errorln(err)
			return
		}

		switch TClientMsgType(vClientMsg.Type()) {
		case CClientMsgType_ToUser:
			this.M_SClientUserMgr.SendToUser(vClientMsg.UserId(), this.M_byMsgBody)
		default:
			GSLog.Errorln("unknown msg type ", vClientMsg.Type())
			return
		}
	}
}
