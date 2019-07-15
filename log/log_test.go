package log

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	const expected = "level=\"info\" message=\"test_msg\""
	writer := &fakeWriter{}
	std = &Logger{
		adapter: NewAdapterLog(writer, "", 0),
	}
	Log(LevelInfo, "test_msg")
	assert.Contains(t, writer.content, expected)
}

func TestConfigure(t *testing.T) {
	adapter := NewAdapterTest(func(fields Fields) {

	})
	Configure(adapter)

	assert.NotNil(t, std)
	assert.Equal(t, std.adapter, adapter)
}

func TestLogFunctions(t *testing.T) {
	cases := []struct {
		level   Level
		logFunc func(string, ...Arg)
	}{
		{LevelTrace, Trace},
		{LevelDebug, Debug},
		{LevelInfo, Info},
		{LevelWarn, Warn},
		{LevelError, Error},
		{LevelPanic, Panic},
		{LevelFatal, Fatal},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.level.String(), func(t *testing.T) {
			t.Parallel()
			testLogFunc(t, c.level, c.logFunc)
		})
	}
}

func testLogFunc(t *testing.T, level Level, logFunc func(string, ...Arg)) {
	expected := fmt.Sprintf("level=\"%v\" message=\"test_msg\"", level)
	writer := &fakeWriter{}
	std = &Logger{
		adapter: NewAdapterLog(writer, "", 0),
	}
	logFunc("test_msg")
	assert.Contains(t, writer.content, expected)
}
