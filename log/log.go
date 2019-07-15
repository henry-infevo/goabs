package log

import (
	"log"
	"os"
)

// Arg log argument function
type Arg func(fields Fields)

var std = &Logger{
	adapter: NewAdapterLog(os.Stderr, "", log.LstdFlags),
}

// Configure configure default log for log package
func Configure(adapter ILogAdapter) {
	std = &Logger{adapter: adapter}
}

// Log log message with level and args
// There are other shorter syntax for specific level
// like: Trace, Debug, Info, Warn, Error, Panic, Fatal
func Log(level Level, msg string, args ...Arg) {
	std.Log(level, msg, args...)
}

// Trace log message with "trace" level and arguments
func Trace(msg string, args ...Arg) {
	std.Trace(msg, args...)
}

// Debug log message with debug level and arguments
func Debug(msg string, args ...Arg) {
	std.Debug(msg, args...)
}

// Info log message with Info level and arguments
func Info(msg string, args ...Arg) {
	std.Info(msg, args...)
}

// Warn log message with "warn" level and arguments
func Warn(msg string, args ...Arg) {
	std.Warn(msg, args...)
}

// Error log message with "error" level
// use WithError to attach the error
func Error(msg string, args ...Arg) {
	std.Error(msg, args...)
}

// Panic log message with "panic" level and arguments
// use WithError to attach the error
// some log engine may call panic function after log the message
func Panic(msg string, args ...Arg) {
	std.Panic(msg, args...)
}

// Fatal log message with "fatal" level and arguments
// use WithError to attach the error
// some log engine may call os.Exit function after log the message
func Fatal(msg string, args ...Arg) {
	std.Fatal(msg, args...)
}
