module github.com/jkaveri/goabs/log/adapter-logrus

go 1.12

replace github.com/jkaveri/goabs/log => ../../log

require (
	github.com/jkaveri/goabs/log v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
)
