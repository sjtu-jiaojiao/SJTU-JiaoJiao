package utils

import (
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
