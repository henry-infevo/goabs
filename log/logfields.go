package log

import (
	"fmt"
	"strings"
)

const (
	FieldLevel   = "level"
	FieldMessage = "message"
	FieldError   = "error"
)

type Fields map[string]interface{}

func (d Fields) Level() LogLevel {
	if val, ok := d[FieldLevel]; ok {
		return val.(LogLevel)
	}
	return LevelTrace
}

func (d Fields) Message() string {
	if val, ok := d[FieldMessage]; ok {
		return val.(string)
	}
	return ""
}

func (d Fields) String() string {
	var sb strings.Builder
	first := true
	for key, val := range d {
		if !first {
			sb.WriteRune(' ')
		} else {
			first = false
		}
		sb.WriteString(fmt.Sprintf("%s=\"%v\"", key, getString(val)))
	}
	return sb.String()
}

func getString(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%d", val)
	case string:
		return escapseDoubleQuote(val.(string))
	}

	if s, ok := val.(stringable); ok {
		return escapseDoubleQuote(s.String())
	}
	return ""
}

type stringable interface {
	String() string
}

func escapseDoubleQuote(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == '"' {
			sb.WriteRune('\\')
		}
		sb.WriteRune(c)
	}
	return sb.String()
}
