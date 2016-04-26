package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"strconv"
)

// IUser
type SServerUser struct {
	*SUser
	M_SServerUserMgr *SServerUserMgr
	M_ServerType     uint16
}

func NewServerUser(vSServerUserMgr *SServerUserMgr) (*SServerUser, error) {
	GSLog.Debugln("newServerUser")
	this := &SServerUser{}

	vSUser, err := NewUser()
	if err != nil {
		return nil, err
	}
	this.SUser = vSUser
	this.M_SServerUserMgr = vSServerUserMgr

	this.M_byMsgHeader = make([]byte, bsn_msg.CSMsgHeaderGateServer_Size)

	return this, nil
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) Run() {
	go this.runImp()
}

func (this *SServerUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		this.UserMgr().DelUser(this)
	}()

	vSClientUserMgr := this.UserMgr().Gate().GetClientMgr()
	var vMsg bsn_msg.SMsgHeaderGateServer
	for !this.M_bClose {
		err := this.ReadMsgHeader()
		if err != nil {
			GSLog.Errorln(err)
			return
		}
		vMsg.DeSerialize(this.M_byMsgHeader)

		vLen := vMsg.Len()
		this.M_byMsgBody = make([]byte, vLen)
		if vLen > 0 {
			err = this.ReadMsgBody()
			if err != nil {
				GSLog.Errorln(err)
				return
			}
		}

		switch vMsg.ServerMsgType() {
		case bsn_msg.CServerMsgType_Reg:
			err = this.msg_reg(&vMsg)
		case bsn_msg.CServerMsgType_ToUser:
			err = vSClientUserMgr.SendMsgToUser(vMsg.UserId(), this.M_byMsgBody)
		default:
			err = errors.New("unknown msg type " + strconv.Itoa(int(vMsg.Type())))
		}

		if err != nil {
			GSLog.Errorln("process msg err: ", vMsg.Type(), err)
		}
	}
}

func (this *SServerUser) msg_reg(vMsg *bsn_msg.SMsgHeaderGateServer) error {
	vServerType := uint8(vMsg.GroupId())
	vServerId := uint8(vMsg.UserId())
	vUserId := bsn_common.MakeGateUserId(vServerType, vServerId)
	GSLog.Debugf("reg server vServerType=%v vServerId=%v vUserId=%v this.Id()=%v \n", vServerType, vServerId, vUserId, this.Id())

	if this.Id() != 0 {
		return errors.New("had reg with id " + strconv.Itoa(int(this.Id())))
	}

	this.SetId(vUserId)
	this.UserMgr().AddUser(this)

	return nil
}
