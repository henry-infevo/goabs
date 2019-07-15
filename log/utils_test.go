package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logStringToMap(t *testing.T) {
	cases := []struct {
		name     string
		log      string
		expected map[string]string
	}{
		{
			name: "happy_case",
			log:  "key=\"val1\"",
			expected: map[string]string{
				"key": "val1",
			},
		},
		{
			name: "happy_case_2",
			log:  "key=\"val1\" key2=\"val2\"",
			expected: map[string]string{
				"key":  "val1",
				"key2": "val2",
			},
		},
		{
			name:     "missing_space",
			log:      "key=\"val1\"key2=\"val2\"",
			expected: map[string]string{},
		},
		{
			name: "missing_double_quote",
			log:  "key=val1 key2=\"val2\"",
			expected: map[string]string{
				"key":  "al",
				"key2": "val2",
			},
		},
	}
	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			m := logStringToMap(c.log)
			assert.Equal(t, c.expected, m)
		})
	}
}
