package utils

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func logLoad() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
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
	if e != nil && e != true {
		s := e
		if v != nil {
			s = fmt.Sprintf(v[0].(string), v[1:]...)
		}
		Error(s)
		panic(s)
	}
}
