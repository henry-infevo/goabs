package log

import (
	"fmt"
)

type LogLevel int
type Arg = func(fields Fields)

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var std ILogger = &FmtLogger{}

func Configure(logger ILogger) {
	std = logger
}

func Log(level LogLevel, msg string, args ...Arg) {
	data := Fields{}
	for _, arg := range args {
		arg(data)
	}
	data[FieldLevel] = level
	data[FieldMessage] = msg
	std.Log(data)
}

// Info log with info level
func Info(msg string, args ...Arg) {
	Log(LevelInfo, msg, args...)
}

func Infof(msg string, args ...interface{}) {
	fArgs := make([]Arg, len(args))
	i := -1
	for _, arg := range args {
		if _, ok := arg.(Arg); ok {
			break
		}
		i++
	}
	if i > -1 {
		msg = fmt.Sprintf(msg, args[0:i])
	}
	Log(LevelInfo, msg)
}
