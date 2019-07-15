package log

// ILogAdapter a log adapter is a connector between `goabs/log` with a log engine
// for example: If you want logging with `logrus` you can implement `ILogAdapter`
// that use `logrus` API to write log based on `log.Fields`
type ILogAdapter interface {
	Log(fields Fields)
}
