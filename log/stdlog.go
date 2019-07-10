package log

import "fmt"

type FmtLogger struct{}

func (*FmtLogger) Log(fields LogFields) {
	msg := fields.String()
	fmt.Println(msg)
}
