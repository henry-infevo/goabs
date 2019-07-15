package log

import (
	"fmt"
	"strings"
)

func getString(val interface{}) string {
	type stringable interface {
		String() string
	}

	switch s := val.(type) {
	case int:
		return fmt.Sprintf("%d", s)
	case string:
		return escapseDoubleQuote(s)
	}

	if s, ok := val.(stringable); ok {
		return escapseDoubleQuote(s.String())
	}
	return ""
}

func escapseDoubleQuote(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == '"' {
			// nolint
			sb.WriteRune('\\')
		}
		// nolint
		sb.WriteRune(c)
	}
	return sb.String()
}

//  nolint
func writeLogItem(sb *strings.Builder, key, val string) {
	sb.WriteString(key)
	sb.WriteRune('=')
	sb.WriteRune('"')
	sb.WriteString(val)
	sb.WriteRune('"')
}

func logStringToMap(str string) map[string]string {
	m := map[string]string{}
	keyPairs := strings.Split(str, " ")
	for _, keyPair := range keyPairs {
		if keyPair == "" {
			continue
		}
		keyval := strings.Split(keyPair, "=")
		if len(keyval) == 2 {
			val := keyval[1]
			if len(val) > 2 {
				val = val[1 : len(val)-1]
			}
			m[keyval[0]] = val
		}
	}
	return m
}
