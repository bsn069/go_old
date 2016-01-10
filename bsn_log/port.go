/*
Package input.

It is generated from these files:
	t.proto

It has these top-level messages:
	Test
*/
package bsn_log

import (
	"fmt"
)

func Print(v ...interface{}) {
	strInfo := fmt.Sprint(v...)
	output(strInfo)
}

func Println(v ...interface{}) {
	strInfo := fmt.Sprint(v...) + "\n"
	output(strInfo)
}

func Printf(format string, v ...interface{}) {
	strInfo := fmt.Sprintf(format, v...)
	output(strInfo)
}
