package bsn_net

import (
	"github.com/bsn069/go/bsn_log"
	"sync"
)

func init() {
	gSTCPRecverPool = sync.Pool{
		New: func() interface{} {
			return &STCPRecver{}
		},
	}

	GSLog = bsn_log.GSLog
}
