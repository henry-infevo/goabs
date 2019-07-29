module github.com/jkaveri/goabs/examples/log-logrus

go 1.12

replace (
	github.com/jkaveri/goabs/log => ../../log/
	github.com/jkaveri/goabs/log/adapter-logrus => ../../log/adapter-logrus/
)

require (
	github.com/jkaveri/goabs/log v1.0.0
	github.com/jkaveri/goabs/log/adapter-logrus v1.0.0
)
