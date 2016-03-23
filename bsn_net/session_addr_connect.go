package bsn_net

import (
	"errors"
	"net"
)

type SSessionAddrConnect struct {
	SSessionAddr
}

func (this *SSessionAddr) Connect() (err error) {
	if "" == this.Addr() {
		return errors.New("no addr")
	}

	if this.Conn() != nil {
		return errors.New("had connect")
	}

	vConn, err := net.Dial("tcp", this.Addr())
	if err != nil {
		return err
	}
	this.SetConn(vConn)

	return nil
}
