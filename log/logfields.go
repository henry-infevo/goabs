package log

import (
	"fmt"
	"strings"
)

const (
	// FieldKeyLevel key of log level field
	FieldKeyLevel = "level"
	// FieldKeyMessage key of log message field
	FieldKeyMessage = "message"
	// FieldKeyError key of log error field
	FieldKeyError = "error"
	// FieldKeyFormatArgs key of format arguments field
	// format arguments is a list of arguments that can
	// be passed into the `fmt.Sprintf`
	FieldKeyFormatArgs = "@@fargs"
)

// Fields fields contains all logging data
type Fields map[string]interface{}

// Level extract level in fields
// return `LevelNone` if there are no level was set or level type
func (d Fields) Level() Level {
	if val, ok := d[FieldKeyLevel]; ok {
		if l, ok := val.(Level); ok {
			return l
		}
	}
	return LevelNone
}

// Message extract log message in fields
// return empty string if log message haven't set yet.
// return formatted string (like `fmt.Sprintf`) if the message is format string
// and user has attached the format arguments by using `WithFormatArgs` method
func (d Fields) Message() string {
	val, ok := d[FieldKeyMessage]
	if !ok {
		return ""
	}
	s, ok := val.(string)
	if !ok {
		return ""
	}
	val, ok = d[FieldKeyFormatArgs]
	if !ok {
		return s
	}

	if args, ok := val.([]interface{}); ok {
		return fmt.Sprintf(s, args...)
	}

	return s
}

// Error extract error from Fields
// return nil if not exist
func (d Fields) Error() error {
	if val, ok := d[FieldKeyError]; ok {
		if err, ok := val.(error); ok {
			return err
		}
	}
	return nil
}

// Rest return all log data in Fields except "known-fields"
// known fields are: Error, Message, Level, FormatArgs
func (d Fields) Rest() Fields {
	rest := Fields{}
	for key, val := range d {
		rest[key] = val
	}
	rest.DeleteAllKnowFields()
	return rest
}

// String serialize the Fields to a string
// return a string in this format:
// ```text
// [level="%s"][ message="%s"][ error="%s"][...[ key="%s"]
// ```
// log message will use the `fields.Message()` method.
func (d Fields) String() string {
	var sb strings.Builder
	str := d.Level().String()
	first := true
	if str != "" {
		first = false
		writeLogItem(&sb, FieldKeyLevel, str)
	}

	str = d.Message()
	if str != "" {
		if !first {
			// nolint
			sb.WriteRune(' ')
		}
		first = false
		writeLogItem(&sb, FieldKeyMessage, str)
	}

	if err := d.Error(); err != nil {
		if !first {
			// nolint
			sb.WriteRune(' ')
		}
		first = false
		writeLogItem(&sb, FieldKeyError, err.Error())
	}
	for key, val := range d {
		// ignore known fields
		if d.IsKnowFieldKey(key) {
			continue
		}
		if !first {
			// nolint
			sb.WriteRune(' ')
		}
		first = false
		writeLogItem(&sb, key, getString(val))
	}
	return sb.String()
}

// IsKnowFieldKey check if field key is know fields
// Know keys are:
// - FieldKeyFormatArgs
// - FieldKeyMessage
// - FieldKeyError
// - FieldKeyLevel
func (d Fields) IsKnowFieldKey(key string) bool {
	return key == FieldKeyFormatArgs ||
		key == FieldKeyMessage ||
		key == FieldKeyError ||
		key == FieldKeyLevel
}

// DeleteAllKnowFields delete all fields which are knew fields
// Know keys are:
// - FieldKeyFormatArgs
// - FieldKeyMessage
// - FieldKeyError
// - FieldKeyLevel
func (d Fields) DeleteAllKnowFields() {
	delete(d, FieldKeyLevel)
	delete(d, FieldKeyError)
	delete(d, FieldKeyMessage)
	delete(d, FieldKeyFormatArgs)
}
