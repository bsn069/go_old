package bsn_gate

import (
	// "strconv"
	"testing"
)

func TestBase(t *testing.T) {
	gate := NewGate()
	serverMgr := gate.GetServerMgr()
	serverMgr.SetListenPort(40000)
	serverMgr.Listen()
}
