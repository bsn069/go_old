package bsn_gate_config

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
)

type SServerUser struct {
	*bsn_net.SNetConnecter
	M_SServerUserMgr *SServerUserMgr
	M_SClientUserMgr *SClientUserMgr

	M_byRecvBuff              []byte
	M_bySMsgHeaderServer2Gate []byte
	M_by2GateMsg              []byte
	M_by2ClientMsg            []byte
	M_SMsgHeaderServer2Gate   bsn_msg.SMsgHeaderServer2Gate
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr, strAddr string) (*SServerUser, error) {
	GSLog.Debugln("NewSServerUser()")

	this := &SServerUser{
		M_SServerUserMgr:          vSServerUserMgr,
		M_SClientUserMgr:          vSServerUserMgr.App().GetClientMgr(),
		M_bySMsgHeaderServer2Gate: make([]byte, bsn_msg.CSMsgHeaderServe2Gater_Size),
		M_byRecvBuff:              make([]byte, 4),
	}
	this.SNetConnecter, _ = bsn_net.NewNetConnecter(this)
	this.SetAddr(strAddr)

	return this, nil
}

func (this *SServerUser) NetConnecterImpRun() error {
	GSLog.Debugln("NetConnecterImpRun")

	err := this.Recv(this.M_bySMsgHeaderServer2Gate)
	if err != nil {
		return err
	}
	GSLog.Debugln("recv this.M_bySMsgHeaderServer2Gate= ", this.M_bySMsgHeaderServer2Gate)

	this.M_SMsgHeaderServer2Gate.DeSerialize(this.M_bySMsgHeaderServer2Gate)
	GSLog.Debugln("recv this.M_SMsgHeaderServer2Gate= ", this.M_SMsgHeaderServer2Gate)

	vTotalLen := int(this.M_SMsgHeaderServer2Gate.Len()) + int(this.M_SMsgHeaderServer2Gate.ServerMsgLen())
	if vTotalLen > cap(this.M_byRecvBuff) {
		// realloc recv buffer
		this.M_byRecvBuff = make([]byte, vTotalLen)
	}
	this.M_byRecvBuff = this.M_byRecvBuff[0:vTotalLen]

	GSLog.Debugln("read byMsgBody")
	err = this.Recv(this.M_byRecvBuff)
	if err != nil {
		return err
	}
	GSLog.Debugln("recv this.M_byRecvBuff= ", this.M_byRecvBuff)

	this.M_by2GateMsg = this.M_byRecvBuff[0:this.M_SMsgHeaderServer2Gate.Len()]
	this.M_by2ClientMsg = this.M_byRecvBuff[this.M_SMsgHeaderServer2Gate.Len():vTotalLen]
	err = this.procMsg()
	if err != nil {
		return err
	}

	return nil
}

func (this *SServerUser) NetConnecterImpOnClose() error {
	GSLog.Debugln("NetConnecterImpOnClose")
	return nil
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) ShowInfo() {
}

func (this *SServerUser) procMsg() error {
	GSLog.Debugln("this.M_SMsgHeaderServer2Gate= ", this.M_SMsgHeaderServer2Gate)
	GSLog.Debugln("this.M_by2GateMsg= ", this.M_by2GateMsg)
	GSLog.Debugln("this.M_by2ClientMsg= ", this.M_by2ClientMsg)
	return nil
}
