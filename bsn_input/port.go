/*
Package bsn_input.
*/
package bsn_input

type IInput interface {
	Reg(strMod string) (chan []string, error)
}

func Instance() IInput {
	return instance()
}

// call in bin file, don`t in lib will block
func Run() {
	run()
}
