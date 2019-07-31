package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchNumberLogFields number of fields that are generated while benchmarking
var BenchNumberLogFields = 50

func TestFields_Message(t *testing.T) {

	type testCase struct {
		name       string
		msg        interface{}
		formatArgs interface{}
		expected   string
	}

	cases := []*testCase{
		{
			name:     "happy_case",
			msg:      "my_msg",
			expected: "my_msg",
		},
		{
			name:     "int_val",
			msg:      10,
			expected: "",
		},
		{
			name:     "nil",
			msg:      nil,
			expected: "",
		},
		{
			name:     "error",
			msg:      errors.New("test val"),
			expected: "",
		},
		{
			name:       "format_msg",
			msg:        "my num: %d, my str: %s",
			formatArgs: []interface{}{10, "a"},
			expected:   "my num: 10, my str: a",
		},
		{
			name:       "format_msg",
			msg:        "my num: %d",
			formatArgs: []int{10},
			expected:   "my num: %d",
		},
	}

	testFunc := func(c *testCase) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			fields := Fields{
				FieldKeyMessage:    c.msg,
				FieldKeyFormatArgs: c.formatArgs,
			}

			result := fields.Message()

			assert.Equal(t, c.expected, result)
		}
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, testFunc(c))
	}
}

func BenchmarkFields_Message(b *testing.B) {
	fields := generateTestFields(BenchNumberLogFields)
	fields[FieldKeyMessage] = "msg"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fields.Message()
	}
}

func TestFields_Error(t *testing.T) {
	type testCase struct {
		name     string
		error    interface{}
		expected error
	}
	expectedErr := errors.New("my_msg")
	cases := []*testCase{
		{
			name:     "happy_case",
			error:    expectedErr,
			expected: expectedErr,
		},
		{
			name:     "int_val",
			error:    10,
			expected: nil,
		},
		{
			name:     "nil",
			error:    nil,
			expected: nil,
		},
	}

	testFunc := func(c *testCase) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			fields := Fields{
				FieldKeyError: c.error,
			}

			result := fields.Error()

			assert.Equal(t, c.expected, result)
		}
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, testFunc(c))
	}
}

func BenchmarkFields_Error(b *testing.B) {
	fields := generateTestFields(BenchNumberLogFields)
	fields[FieldKeyError] = errors.New("error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// nolint
		fields.Error()
	}
}

func TestFields_String(t *testing.T) {
	type testCase struct {
		name     string
		fields   Fields
		expected map[string]string
	}
	cases := []*testCase{
		{
			name: "happy_case",
			fields: Fields{
				"s":             "my_val",
				"i":             10,
				FieldKeyLevel:   LevelDebug,
				FieldKeyMessage: "my_msg",
				FieldKeyError:   errors.New("my_error"),
			},
			expected: map[string]string{
				"s":             "my_val",
				"i":             "10",
				FieldKeyLevel:   LevelDebug.String(),
				FieldKeyMessage: "my_msg",
				FieldKeyError:   "my_error",
			},
		},
		{
			name: "without_level",
			fields: Fields{
				"s":             "my_val",
				"i":             10,
				FieldKeyMessage: "my_msg",
				FieldKeyError:   errors.New("my_error"),
			},
			expected: map[string]string{
				"s":             "my_val",
				"i":             "10",
				FieldKeyMessage: "my_msg",
				FieldKeyError:   "my_error",
			},
		},
		{
			name: "without_msg",
			fields: Fields{
				"s":           "my_val",
				"i":           10,
				FieldKeyLevel: LevelDebug,
				FieldKeyError: errors.New("my_error"),
			},
			expected: map[string]string{
				"s":           "my_val",
				"i":           "10",
				FieldKeyLevel: LevelDebug.String(),
				FieldKeyError: "my_error",
			},
		},
		{
			name: "without_error",
			fields: Fields{
				"s":             "my_val",
				"i":             10,
				FieldKeyLevel:   LevelDebug,
				FieldKeyMessage: "my_msg",
			},
			expected: map[string]string{
				"s":             "my_val",
				"i":             "10",
				FieldKeyLevel:   LevelDebug.String(),
				FieldKeyMessage: "my_msg",
			},
		},
		{
			name: "without_error_msg_level",
			fields: Fields{
				"s": "my_val",
				"i": 10,
			},
			expected: map[string]string{
				"s": "my_val",
				"i": "10",
			},
		},
		{
			name:     "empty",
			fields:   Fields{},
			expected: map[string]string{},
		},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := c.fields.String()
			if result != "" {
				assert.True(t, result[0] != ' ')
				assert.True(t, result[len(result)-1] != ' ')
			}
			m := logStringToMap(result)
			assert.Equal(t, c.expected, m)
		})
	}
}

func BenchmarkFields_String(b *testing.B) {
	fields := generateTestFields(BenchNumberLogFields)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fields.String()
	}
}

func BenchmarkFields_Rest(b *testing.B) {
	fields := generateTestFields(BenchNumberLogFields)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fields.Rest()
	}
}

func generateTestFields(num int) Fields {
	fields := Fields{}
	var key string
	var val interface{}
	for i := 0; i < num-1; i++ {
		key = fmt.Sprintf("%d", i)
		val = i
		fields[key] = val
	}
	return fields
}
