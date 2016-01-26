package bsn_gate

import (
	// "strconv"
	"testing"
)

func TestBase(t *testing.T) {
	gate := NewGate()
	clientMgr := gate.GetClientMgr()
	clientMgr.SetListenPort(50000)
	clientMgr.Listen()

	serverMgr := gate.GetServerMgr()
	serverMgr.SetListenPort(40000)
	serverMgr.Listen()
}
