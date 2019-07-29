package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_Log(t *testing.T) {
	type testArgs struct {
		level Level
		msg   string
		args  []Arg
	}
	type testWant struct {
		fields Fields
	}
	type testCase struct {
		name string
		args *testArgs
		want *testWant
	}

	testFunc := func(c *testCase) func(t *testing.T) {
		return func(t *testing.T) {
			logFunc := func(fields Fields) {
				assert.Equal(t, c.want.fields, fields)
			}
			logger := NewLogger(NewAdapterTest(logFunc))

			logger.Log(c.args.level, c.args.msg, c.args.args...)
		}
	}

	cases := []*testCase{
		{
			name: "should_success",
			args: &testArgs{
				level: LevelInfo,
				msg:   "test",
				args:  nil,
			},
			want: &testWant{
				fields: map[string]interface{}{
					FieldKeyMessage: "test",
					FieldKeyLevel:   LevelInfo,
				},
			},
		},
		{
			name: "should_combine_fields",
			args: &testArgs{
				level: LevelInfo,
				msg:   "test",
				args: []Arg{
					func(fields Fields) {
						fields["test_arg"] = "test_arg"
						fields["a_number"] = 10
					},
				},
			},
			want: &testWant{
				fields: map[string]interface{}{
					FieldKeyMessage: "test",
					FieldKeyLevel:   LevelInfo,
					"test_arg":      "test_arg",
					"a_number":      10,
				},
			},
		},
		{
			name: "should_not_overwrite_default_fields",
			args: &testArgs{
				level: LevelInfo,
				msg:   "test",
				args: []Arg{
					func(fields Fields) {
						fields[FieldKeyLevel] = LevelDebug
						fields[FieldKeyMessage] = "no_test_message"
					},
				},
			},
			want: &testWant{
				fields: map[string]interface{}{
					FieldKeyMessage: "test",
					FieldKeyLevel:   LevelInfo,
				},
			},
		},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, testFunc(c))
	}
}

func TestLogger_Trace(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelTrace, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Trace(msg)
}

func TestLogger_Debug(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelDebug, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Debug(msg)
}

func TestLogger_Info(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelInfo, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Info(msg)
}

func TestLogger_Warn(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelWarn, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Warn(msg)
}

func TestLogger_Error(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelError, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Error(msg)
}

func TestLogger_Panic(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelPanic, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Panic(msg)
}

func TestLogger_Fatal(t *testing.T) {
	const msg = "message"
	logger := NewLogger(NewAdapterTest(func(fields Fields) {
		assert.Equal(t, LevelFatal, fields.Level())
		assert.Equal(t, msg, fields.Message())
	}))

	logger.Fatal(msg)
}
