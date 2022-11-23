package cast

type Logger interface {
	Printf(format string, v ...interface{})
}

var logger Logger = &noopLogger{}

type noopLogger struct{}

func (n *noopLogger) Printf(format string, v ...interface{}) {}

func SetLogger(l Logger) {
	logger = l
}
