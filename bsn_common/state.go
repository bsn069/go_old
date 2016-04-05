package bsn_common

import (
	"sync/atomic"
)

type TState int32

const (
	CState_Idle TState = iota
	CState_PrepareRuning
	CState_Runing
	CState_Stoping
	CState_StopFromStopListen
	CState_StopFromAcceptClose
	CState_Op
)

func (this TState) String() string {
	switch this {
	case CState_Idle:
		return "Idle"
	case CState_Runing:
		return "Runing"
	case CState_Stoping:
		return "Stoping"
	}
	return "unknown"
}

type SState struct {
	M_TState TState
}

func NewSState() *SState {
	this := &SState{
		M_TState: CState_Idle,
	}
	return this
}

func (this *SState) Change(from, to TState) bool {
	return atomic.CompareAndSwapInt32((*int32)(&this.M_TState), int32(from), int32(to))
}

func (this *SState) Is(vTState TState) bool {
	return atomic.LoadInt32((*int32)(&this.M_TState)) == int32(vTState)
}

func (this *SState) Set(vTState TState) {
	atomic.StoreInt32((*int32)(&this.M_TState), int32(vTState))
}
