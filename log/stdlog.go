package log

import "fmt"

type FmtLogger struct{}

func (*FmtLogger) Log(fields Fields) {
	msg := fields.String()
	fmt.Println(msg)
}
