# Logging

This package provide methods that help in logging you can change the log engine by changing the log adapter which implements the `ILogAdapter` interface

## Design

![log package design](./design.png)

## Usage

### Default

You can start using log package with default configuration which will use `AdapterLog`  as an adapter which uses "log" package as log engine

```go
func man() {
	// simple log with level "info"
	log.Info("this is an log message")
	// log with error
	log.Error(
		"this is an error log message", 
		log.WithError(errors.New("simple eror")),
	)
	// log with custom field
	log.Trace(
		"this is an trace message with username",
		log.WithField("username", "jkaveri"),
		log.WithField("repo", "github.com/jkaveri/goabs"),
	)
}
```

### Use another log engine

You can use any log engine which you like and familiar with. Just implement the `ILogAdapter`

```go
type ILogAdapter interface {
	Log(fields Fields)
}
```

When you implemented the adapter you can configure log package to use your adapter:

```go
log.Configure(mylogadapter.NewAdapter())
```

### Examples

- [Use with logrus](../examples/log-logrus)