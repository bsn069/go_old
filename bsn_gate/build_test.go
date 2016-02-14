package bsn_gate

import (
	// "strconv"
	// "net"
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	gate, err := NewGate()
	if err != nil {
		t.Fatal(err)
	}
	gate.GetServerMgr().SetListenAddr("localhost:40000")
	gate.GetClientMgr().SetListenAddr("localhost:50000")
	gate.Listen()
	// serverMgr := gate.GetServerMgr()
	// serverMgr.SetListenPort(40000)

	// clientMgr := gate.GetClientMgr()
	// clientMgr.SetListenPort(50000)

	// for i := 0; i < 100; i++ {
	// 	gate.Listen()

	// 	for j := 0; j < 2; j++ {
	// 		iConn, err := net.Dial("tcp", "127.0.0.1:40000")
	// 		if err != nil {
	// 			GLog.Errorln(err)
	// 			t.Fatal(err)
	// 		}
	// 		iConn.Close()
	// 	}

	// 	for j := 0; j < 2; j++ {
	// 		iConn, err := net.Dial("tcp", "127.0.0.1:50000")
	// 		if err != nil {
	// 			GLog.Errorln(err)
	// 			t.Fatal(err)
	// 		}
	// 		iConn.Close()
	// 	}
	// 	select {
	// 	case <-time.After(100):
	// 	}
	// 	gate.Close()
	// }

	select {
	case <-time.After(10):
	}
	gate.Close()
	t.Log("close")

	select {
	case <-time.After(10):
	}
	t.Log("end")
}
