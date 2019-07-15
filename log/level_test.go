package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	cases := []struct {
		name      string
		level     Level
		levelName string
	}{
		{"trace", LevelTrace, "trace"},
		{"debug", LevelDebug, "debug"},
		{"info", LevelInfo, "info"},
		{"warn", LevelWarn, "warn"},
		{"error", LevelError, "error"},
		{"panic", LevelPanic, "panic"},
		{"fatal", LevelFatal, "fatal"},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			levelName := c.level.String()

			assert.Equal(t, c.levelName, levelName)
		})
	}
}
