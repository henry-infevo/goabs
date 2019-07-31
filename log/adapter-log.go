package log

import (
	"io"
	"log"
	"os"
)

// AdapterLog will be used as default log adapter
type AdapterLog struct {
	logger *log.Logger
}

// compile time error if AdapterLog not implement ILogAdapter
var _ ILogAdapter = (*AdapterLog)(nil)

// NewAdapterLog create `AdapterLog``. The `out`
// destination to which log data will be written. default: `os.Stderr`
// The `prefix` appears at the beginning of each generated log line.
// The `flag` argument defines the logging properties.
// refer: https://golang.org/pkg/log/#pkg-constants
func NewAdapterLog(out io.Writer, prefix string, flag int) *AdapterLog {
	if out == nil {
		out = os.Stderr
	}
	return &AdapterLog{
		logger: log.New(out, prefix, flag),
	}
}

// Log will use go native `"log"` package to write log from `fields`
func (t *AdapterLog) Log(fields Fields) {
	level := fields.Level()
	msg := fields.String()

	if level == LevelPanic {
		t.logger.Panicln(msg)
		return
	}

	if level == LevelFatal {
		t.logger.Fatalln(msg)
		return
	}

	t.logger.Println(msg)
}
