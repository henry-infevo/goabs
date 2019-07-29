package log

// Logger provide common logging methods
// the logger will call `ILogAdapter.Log` to write actual log item
type Logger struct {
	adapter ILogAdapter
}

// NewLogger create logger instance
func NewLogger(adapter ILogAdapter) *Logger {
	return &Logger{adapter}
}

// Log a generic method that write log with level and msg
// you can pass-in the list of logging arguments to add more information into the log item
// Use some sugar method like: WithField, WithError, WithFormatArgs
// the level and the msg cannot be override by `Arg` function
func (t *Logger) Log(level Level, msg string, args ...Arg) {
	fields := Fields{}
	for _, arg := range args {
		arg(fields)
	}
	fields[FieldKeyLevel] = level
	fields[FieldKeyMessage] = msg
	t.adapter.Log(fields)
}

// Trace log msg with "trace" level
func (t *Logger) Trace(msg string, args ...Arg) {
	t.Log(LevelTrace, msg, args...)
}

// Debug log msg with "debug" level
func (t *Logger) Debug(msg string, args ...Arg) {
	t.Log(LevelDebug, msg, args...)
}

// Info log msg with "info" level
func (t *Logger) Info(msg string, args ...Arg) {
	t.Log(LevelInfo, msg, args...)
}

// Warn log msg with "warn" level
func (t *Logger) Warn(msg string, args ...Arg) {
	t.Log(LevelWarn, msg, args...)
}

// Error log msg with "error" level
// use WithError to attach the error
func (t *Logger) Error(msg string, args ...Arg) {
	t.Log(LevelError, msg, args...)
}

// Panic log msg with "panic" level
// use WithError to attach the error
// some log engine will call `panic` when logging this level
func (t *Logger) Panic(msg string, args ...Arg) {
	t.Log(LevelPanic, msg, args...)
}

// Fatal log msg with "fatal" level
// use WithError to attach the error
// some log engine will call `os.Exit` when logging this level
func (t *Logger) Fatal(msg string, args ...Arg) {
	t.Log(LevelFatal, msg, args...)
}
