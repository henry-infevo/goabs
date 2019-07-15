package log

// Level logging level type
type Level int

const (
	// LevelNone log level that represent as not set level
	LevelNone Level = iota
	// LevelTrace log level trace
	LevelTrace
	// LevelDebug log level debug
	LevelDebug
	// LevelInfo log level info
	LevelInfo
	// LevelWarn log level warn
	LevelWarn
	// LevelError log level error
	LevelError
	// LevelPanic log level panic
	LevelPanic
	// LevelFatal log level fatal
	LevelFatal
)

// String return log level name
func (t Level) String() string {
	switch t {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelPanic:
		return "panic"
	case LevelFatal:
		return "fatal"
	default:
		return ""
	}
}
