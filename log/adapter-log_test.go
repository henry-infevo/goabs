package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeWriter struct {
	content string
}

func (t *fakeWriter) Write(data []byte) (int, error) {
	t.content = string(data)
	return len(data), nil
}

func TestAdapterLog_Log(t *testing.T) {
	cases := []struct {
		name   string
		fields Fields
		flag   int
	}{
		{
			name:   "only_msg",
			fields: Fields{fieldMessage: "msg"},
			flag:   0,
		},
		{
			name: "only_level",
			fields: Fields{
				fieldLevel: LevelInfo,
			},
			flag: 0,
		},
		{
			name: "only_level",
			fields: Fields{
				fieldLevel: LevelInfo,
			},
			flag: 0,
		},
		{
			name: "only_error",
			fields: Fields{
				fieldLevel: LevelError,
			},
			flag: 0,
		},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			writer := &fakeWriter{}
			adapter := NewAdapterLog(writer, "", 0)
			adapter.Log(c.fields)
			content := writer.content
			assert.Contains(t, content, c.fields.String())
		})
	}

}
