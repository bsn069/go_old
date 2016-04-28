package bsn_net

type SSessionAddr struct {
	SSession
	M_strAddr string
}

func NewSSessionAddr() (*SSessionAddr, error) {
	GSLog.Debugln("NewSSessionAddr")

	this := &SSessionAddr{}

	return this, nil
}

func (this *SSessionAddr) SetAddr(strAddr string) error {
	this.M_strAddr = strAddr
	return nil
}

func (this *SSessionAddr) Addr() string {
	return this.M_strAddr
}
