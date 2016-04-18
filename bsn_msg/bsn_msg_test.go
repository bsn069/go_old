package bsn_msg

import (
	// "strconv"
	"github.com/bsn069/go/bsn_common"
	"testing"
	"unsafe"
)

func TestBase(t *testing.T) {

}

func TestMsgHeader(t *testing.T) {
	msgHeader := NewMsgHeader(1, 2)
	if msgHeader.Type() != 1 {
		t.Error(msgHeader.Type())
	}
	if msgHeader.Len() != 2 {
		t.Error(msgHeader.Len())
	}

	if bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)) != CSMsgHeader_Size {
		t.Errorf("bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader))=%v != CSMsgHeader_Size=%v", bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)), CSMsgHeader_Size)
	}

	byDatas := msgHeader.Serialize()
	GSLog.Mustln(byDatas)

	msgHeader = NewMsgHeaderFromBytes(byDatas)
	if msgHeader.Type() != 1 {
		t.Error(msgHeader.Type())
	}
	if msgHeader.Len() != 2 {
		t.Error(msgHeader.Len())
	}
}

func TestMsgHeaderGateClient(t *testing.T) {
	byMsg := make([]byte, 4)
	msgHeader := NewMsgHeaderGateClient(1, 2, bsn_common.TMsgLen(len(byMsg)))

	vUserId := bsn_common.MakeGateUserId(1, 2)
	if msgHeader.UserId() != vUserId {
		t.Error(msgHeader.UserId())
	}
	if msgHeader.Len() != 4 {
		t.Error(msgHeader.Len())
	}

	if bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)) != CSMsgHeaderGateClient_Size {
		t.Errorf("bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader))=%v != CSMsgHeaderGateClient_Size=%v", bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)), CSMsgHeaderGateClient_Size)
	}

	byDatas := msgHeader.Serialize()
	GSLog.Mustln(byDatas)

	msgHeader = NewMsgHeaderGateClientFromBytes(byDatas)

	if msgHeader.UserId() != vUserId {
		t.Error(msgHeader.UserId())
	}
	if msgHeader.Len() != 4 {
		t.Error(msgHeader.Len())
	}
}

func TestMsgHeaderGateServer(t *testing.T) {
	byMsg := make([]byte, 4)
	msgHeader := NewMsgHeaderGateServer(CServerMsgType_ToUser, 1, 2, bsn_common.TMsgLen(len(byMsg)))
	if msgHeader.ServerMsgType() != CServerMsgType_ToUser {
		t.Error(msgHeader.ServerMsgType())
	}
	if msgHeader.GroupId() != 1 {
		t.Error(msgHeader.GroupId())
	}
	if msgHeader.UserId() != 2 {
		t.Error(msgHeader.UserId())
	}
	if msgHeader.Len() != 4 {
		t.Error(msgHeader.Len())
	}

	if bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)) != CSMsgHeaderGateServer_Size {
		t.Errorf("bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader))=%v != CSMsgHeaderGateServer_Size=%v", bsn_common.TMsgLen(unsafe.Sizeof(*msgHeader)), CSMsgHeaderGateServer_Size)
	}

	byDatas := msgHeader.Serialize()
	GSLog.Mustln(byDatas)

	msgHeader = NewMsgHeaderGateServerFromBytes(byDatas)
	if msgHeader.ServerMsgType() != CServerMsgType_ToUser {
		t.Error(msgHeader.ServerMsgType())
	}
	if msgHeader.GroupId() != 1 {
		t.Error(msgHeader.GroupId())
	}
	if msgHeader.UserId() != 2 {
		t.Error(msgHeader.UserId())
	}
	if msgHeader.Len() != 4 {
		t.Error(msgHeader.Len())
	}
}

func TestMsgHeaderServer2Gate(t *testing.T) {
	vMsg := new(SMsgHeaderServer2Gate)
	GSLog.Mustln("vMsg=", vMsg)

	by2GateMsg := make([]byte, 2)
	by2ClientMsg := make([]byte, 5)
	vMsg.FillMsg(CServer2GateMsgType_Test, by2GateMsg, by2ClientMsg)
	GSLog.Mustln("vMsg=", vMsg)

	byMsg := vMsg.Serialize()
	GSLog.Mustln("byMsg=", byMsg)

}

func TestSMsg_Gate2Server_ClientMsg(t *testing.T) {
	vMsg := new(SMsg_Gate2Server_ClientMsg)
	GSLog.Mustln("vMsg=", vMsg)

	vSMsgHeader := new(SMsgHeader)
	vSMsgHeader.Fill(2, 0)
	vMsg.Fill(1, vSMsgHeader, nil)
	GSLog.Mustln("vMsg=", vMsg)
	GSLog.Mustln("vMsg=", vMsg.M_ClientId, vMsg.M_SMsgHeader, vMsg.M_byMsgBody)

	vbyMsg := vMsg.Serialize()
	GSLog.Mustln("vbyMsg=", vbyMsg)

	vSMsgHeader.DeSerialize(vbyMsg)
	GSLog.Mustln("vMsg=", vMsg.M_ClientId, vMsg.M_SMsgHeader, vMsg.M_byMsgBody)

}

// go test -test.bench=".*"
func BenchmarkA(b *testing.B) {

}
