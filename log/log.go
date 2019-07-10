package log

import (
	"fmt"
	"strings"
)

type LogLevel int
type LogFields map[string]interface{}
type Arg = func(fields LogFields)

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	FieldLevel   = "level"
	FieldMessage = "message"
	FieldError   = "error"
)

type ILogger interface {
	Log(fields LogFields)
}

var std ILogger

func init() {
	std = &FmtLogger{}
}

func Configure(logger ILogger) {
	std = logger
}

func Log(args ...Arg) {
	data := LogFields{}
	for _, arg := range args {
		arg(data)
	}
	std.Log(data)
}

func Level(level LogLevel) Arg {
	return func(data LogFields) {
		data[FieldLevel] = level
	}
}
func Message(message string) Arg {
	return func(data LogFields) {
		data[FieldMessage] = message
	}
}
func Fields(fields LogFields) Arg {
	return func(finalFields LogFields) {
		for key, val := range finalFields {
			finalFields[key] = val
		}
	}
}

// func Trace() {
// 	Log(LevelTrace, args...)
// }
// func Debug(args ...interface{}) {
// 	Log(LevelDebug, args...)
// }
// func Info(args ...interface{}) {
// 	Log(LevelInfo, args...)
// }
// func Warn(args ...interface{}) {
// 	Log(LevelWarn, args...)
// }
// func Error(args ...interface{}) {
// 	Log(LevelError, args...)
// }
// func Fatal(args ...interface{}) {
// 	Log(LevelFatal, args...)
// }
func (d LogFields) String() string {
	var sb strings.Builder
	first := true
	for key, val := range d {
		if !first {
			sb.WriteRune(' ')
		} else {
			first = false
		}
		sb.WriteString(fmt.Sprintf("%s=\"%v\"", key, getString(val)))
	}
	return sb.String()
}

func getString(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%d", val)
	case string:
		return escapseDoubleQuote(val.(string))
	}

	if s, ok := val.(stringable); ok {
		return escapseDoubleQuote(s.String())
	}
	return ""
}

type stringable interface {
	String() string
}

func escapseDoubleQuote(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == '"' {
			sb.WriteRune('\\')
		}
		sb.WriteRune(c)
	}
	return sb.String()
}
