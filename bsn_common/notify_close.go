package bsn_common

import (
// "sync/atomic"
)

type SNotifyClose struct {
	M_chanNotifyClose chan bool
	M_chanWaitClose   chan bool
}

func NewSNotifyClose() *SNotifyClose {
	this := &SNotifyClose{}
	this.Reset()
	return this
}

func (this *SNotifyClose) Reset() {
	this.M_chanNotifyClose = make(chan bool, 1)
	this.M_chanWaitClose = make(chan bool, 1)
}

func (this *SNotifyClose) NotifyClose() {
	this.M_chanNotifyClose <- true
}

func (this *SNotifyClose) WaitClose() {
	<-this.M_chanWaitClose
}

func (this *SNotifyClose) Close() {
	this.M_chanWaitClose <- true
}

func (this *SNotifyClose) Clear() {
	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	select {
	case <-this.M_chanWaitClose:
	default:
	}
}
