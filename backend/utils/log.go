package utils

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

// LoadLog load log config
func LoadLog() {
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	if LocalConf.Deploy == "product" {
		logs.Async()
	}
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	logs.Warning(f, v...)
}

// Info logs a message at information level.
func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

/*
LogPanic will log and panic
if e is nil or true, nothing happens, if v is empty, e will be logged when not nil.

example:
e := nil
LogPanic(e)	// nothing
e = true
LogPanic(e,"123%d",4) // nothing
e = "123"
LogPanic(e)	// panic 123
LogPanic(e,"123%d",4)	// panic 1234
*/
func LogPanic(e interface{}, v ...interface{}) {
	r, s := LogContinueS(e, Error, v...)
	if r {
		panic(s)
	}
}

/*
LogContinueS will log but not panic
if e is nil or true, nothing happens, if v is empty, e will be logged when not nil.
return true when trigger log with log info
*/
func LogContinueS(e interface{}, f func(f interface{}, v ...interface{}), v ...interface{}) (bool, string) {
	if e != nil && e != true {
		s := fmt.Sprintf("%v", e)
		if e == false {
			s = "false"
		}
		if v != nil {
			s = fmt.Sprintf(v[0].(string), v[1:]...)
		}
		f(s)
		return true, s
	}
	return false, ""
}

/*
LogContinue will log but not panic
if e is nil or true, nothing happens, if v is empty, e will be logged when not nil.
return true when trigger log
*/
func LogContinue(e interface{}, f func(f interface{}, v ...interface{}), v ...interface{}) bool {
	r, _ := LogContinueS(e, f, v...)
	return r
}
