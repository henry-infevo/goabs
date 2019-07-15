package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithField(t *testing.T) {
	arg := WithField("a", 10)
	fields := Fields{}
	arg(fields)

	assert.Equal(t, 10, fields["a"])
}

func TestWithFields(t *testing.T) {
	expected := Fields{
		"a": 10,
		"b": "c",
	}
	arg := WithFields(expected)
	fields := Fields{}
	arg(fields)
	assert.Equal(t, expected, fields)
}

func TestWithFormatArg(t *testing.T) {
	arg := WithFormatArg("1", 2, true)
	fields := Fields{}
	arg(fields)

	assert.Equal(t, []interface{}{"1", 2, true}, fields[fieldFormatArgs])
}

func TestWithFormatArgAndMessage(t *testing.T) {
	arg := WithFormatArg("1", 2, true)
	fields := Fields{
		fieldMessage: "%v, %v, %v",
	}
	arg(fields)

	assert.Equal(t, []interface{}{"1", 2, true}, fields[fieldFormatArgs])
	assert.Equal(t, "1, 2, true", fields.Message())
}
