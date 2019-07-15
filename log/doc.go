// Package log provide common logging methods and `ILogAdapter` interface that help we
// wrap another log engine package such as `logrus`, `zap` or `zerolog`
// we can archive the flexibility by wrapping the log engine package. In another word
// our application can be easy to switch to another log engine when we want and it should not
// have any impact
package log
