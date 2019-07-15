package log

// AdapterTest is a stub help on testing
type AdapterTest struct {
	logFunc func(fields Fields)
}

// NewAdapterTest create `AdapterTest` by provide a function that help
// inspect the `fields` which were passed by `Log` method
func NewAdapterTest(logFunc func(fields Fields)) ILogAdapter {
	return &AdapterTest{
		logFunc: logFunc,
	}
}

// Log implement `ILogAdapter` interface
// the fields pass-in will be pass to the logFunc which was initiate in the `NewAdapterTest` method
func (t *AdapterTest) Log(fields Fields) {
	t.logFunc(fields)
}
