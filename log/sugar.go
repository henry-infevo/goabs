package log

// WithField attach field into a log item
func WithField(key string, val interface{}) Arg {
	return func(fields Fields) {
		fields[key] = val
	}
}

// WithFields attach multiple fields into a log item
func WithFields(fields Fields) Arg {
	return func(finalFields Fields) {
		for key, val := range fields {
			finalFields[key] = val
		}
	}
}

// WithFormatArg if the message is a format string,
// use this method to pass the format arguments.
// when serialize the log message this format arguments
// and the format string will be passed into the `fmt.Sprintf`
// method
func WithFormatArg(args ...interface{}) Arg {
	return func(fields Fields) {
		fields[FieldKeyFormatArgs] = args
	}
}

// WithError will attach the error into the log item fields
func WithError(err error) Arg {
	return func(fields Fields) {
		fields[FieldKeyError] = err
	}
}
