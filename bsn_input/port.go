/*
Package bsn_input.
*/
package bsn_input

import (
	"github.com/bsn069/go/bsn_log"
)

type IInput interface {
	Reg(strMod string) (chan []string, error)
}

var GLog = bsn_log.New()
var Instance = instance

// call in bin file, don`t in lib will block
var Run = run
