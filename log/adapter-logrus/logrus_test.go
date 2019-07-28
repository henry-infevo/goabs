package logrus

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/jkaveri/goabs/log"
)

func TestAdapterLogrus_Log(t *testing.T) {
	var sb strings.Builder
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(&sb)
	adapter := &AdapterLogrus{
		logger: logrusLogger,
	}

	adapter.Log(log.Fields{
		"data":            "1",
		"num":             10,
		"level":           log.LevelInfo,
		"message":         "test_msg abc",
		log.FieldKeyError: errors.New("some error"),
	})
	str := sb.String()
	assert.Contains(t, str, "data=1")
	assert.Contains(t, str, "num=10")
	assert.Contains(t, str, "level=info")
	assert.Contains(t, str, "msg=\"test_msg abc\"")
	assert.Contains(t, str, "error=\"some error\"")
}

func BenchmarkAdapterLogrus_Log(b *testing.B) {
	var sb strings.Builder
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(&sb)
	adapter := &AdapterLogrus{
		logger: logrusLogger,
	}

	for i := 0; i < b.N; i++ {
		fields := generateLogFields(50)
		b.StartTimer()
		adapter.Log(fields)
		b.StopTimer()
	}
}

func BenchmarkLogrus(b *testing.B) {
	var sb strings.Builder
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(&sb)
	for i := 0; i < b.N; i++ {
		fields := generateLogrusFields(50)
		b.StartTimer()
		logrusLogger.WithFields(fields).Log(logrus.InfoLevel)
		b.StopTimer()
	}
}

func generateLogFields(num int) log.Fields {
	fields := log.Fields{}
	var key string
	var val interface{}
	for i := 0; i < num-1; i++ {
		key = fmt.Sprintf("%d", i)
		val = i
		fields[key] = val
	}
	return fields
}

func generateLogrusFields(num int) logrus.Fields {
	fields := logrus.Fields{}
	var key string
	var val interface{}
	for i := 0; i < num-1; i++ {
		key = fmt.Sprintf("%d", i)
		val = i
		fields[key] = val
	}
	return fields
}
