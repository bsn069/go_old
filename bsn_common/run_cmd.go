package bsn_common

/*
call func on one routine
*/
import (
	"errors"
)

type TRunCmdFuncOnQuit func() error
type SRunCmd struct {
	M_chanFunc          chan func() error
	M_str2Func          map[string]func() error
	M_bQuit             bool
	M_TRunCmdFuncOnQuit TRunCmdFuncOnQuit
	M_ILog              ILog
}

func NewSRunCmd(vTRunCmdFuncOnQuit TRunCmdFuncOnQuit, vILog ILog) (this *SRunCmd, err error) {
	vILog.Debugln("NewSRunCmd()")
	this = &SRunCmd{
		M_chanFunc:          make(chan func() error),
		M_str2Func:          make(map[string]func() error),
		M_TRunCmdFuncOnQuit: vTRunCmdFuncOnQuit,
		M_ILog:              vILog,
	}

	this.RegCmd("help", this.help)
	this.RegCmd("quit", this.quit)

	go this.mainLoop()
	return this, nil
}

func (this *SRunCmd) Do(vstrCmd string) (err error) {
	defer func() {
		if err != nil {
			this.M_ILog.Errorln(err)
		}
	}()

	if vFunc, ok := this.M_str2Func[vstrCmd]; ok {
		this.M_chanFunc <- vFunc
	} else {
		this.Do("help")
		err = errors.New("not found cmd")
	}
	return
}

func (this *SRunCmd) RegCmd(vstrCmd string, vFunc func() error) {
	this.M_str2Func[vstrCmd] = vFunc
	return
}

func (this *SRunCmd) mainLoop() (err error) {
	for !this.M_bQuit {
		vFunc := <-this.M_chanFunc
		vFunc()
	}
	this.M_ILog.Mustln("quit main loop")

	this.M_TRunCmdFuncOnQuit()
	return
}

func (this *SRunCmd) help() (err error) {
	for vstrCmd, _ := range this.M_str2Func {
		this.M_ILog.Mustln(vstrCmd)
	}
	return
}

func (this *SRunCmd) quit() (err error) {
	this.M_bQuit = true
	return
}
