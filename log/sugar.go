package log

func WithField(key string, val interface{}) Arg {
	return func(fields Fields) {
		fields[key] = val
	}
}

func WithFields(fields Fields) Arg {
	return func(finalFields Fields) {
		for key, val := range finalFields {
			finalFields[key] = val
		}
	}
}
