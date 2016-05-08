package bsn_common

type ILog interface {
	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})
	Must(v ...interface{})
	Mustln(v ...interface{})
	Mustf(format string, v ...interface{})
}
