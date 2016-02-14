package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

// IUser
type SServerUser struct {
	*SUser
	M_SServerUserMgr *SServerUserMgr
}

func newServerUser(vSServerUserMgr *SServerUserMgr) (*SServerUser, error) {
	GLog.Debugln("newServerUser")
	this := &SServerUser{}

	vSUser, err := newUser()
	if err != nil {
		return nil, err
	}
	this.SUser = vSUser
	this.M_SServerUserMgr = vSServerUserMgr

	this.M_byMsgHeader = make([]byte, CSServerMsg_Size)

	return this, nil
}

func (this *SServerUser) Run() {
	go this.runImp()
}

func (this *SServerUser) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.M_Conn.Close()
	return nil
}

func (this *SServerUser) Send(byData []byte) error {
	this.M_Conn.Write(byData)
	return nil
}

func (this *SServerUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		this.M_SServerUserMgr.DelUser(this)
	}()

	var vServerMsg SServerMsg
	for !this.M_bClose {
		err := this.ReadMsgHeader()
		if err != nil {
			GLog.Errorln(err)
			return
		}
		vServerMsg.DeSerialize(this.M_byMsgHeader)

		vLen := vServerMsg.Len()
		if vLen < bsn_msg.CSMsgHeader_Size {
			GLog.Errorln("too short")
			return
		}

		this.M_byMsgBody = make([]byte, vLen)
		err = this.ReadMsgBody()
		if err != nil {
			GLog.Errorln(err)
			return
		}

		switch TServerMsgType(vServerMsg.Type()) {
		case CServerMsgType_ToUser:
			this.M_SServerUserMgr.SendToUser(vServerMsg.UserId(), this.M_byMsgBody)
		default:
			GLog.Errorln("unknown msg type ", vServerMsg.Type())
			return
		}
	}
}
