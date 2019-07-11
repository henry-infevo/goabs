package log

type ILogger interface {
	Log(fields Fields)
}
