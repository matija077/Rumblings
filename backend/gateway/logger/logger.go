package logger

import "os"

type LoggerTypes string

const (
	Basic   LoggerTypes = "basic"
	Simple  LoggerTypes = "simple"
	Advance LoggerTypes = "advance"
)

type LoggerEngineI interface {
	createFile(name string) (*os.File, error)
}

type loggerEngineT struct {
}

func (loggerEngine *loggerEngineT) createFile(name string) (*os.File, error) {
	//f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//	var file, err = os.Create()

	return nil, nil
}

type LoggerI interface {
	Add()
	Search()
}

var loggerEngine LoggerEngineI

func init() {
	loggerEngine = &loggerEngineT{}
}

func createLogger(loggerType LoggerTypes) LoggerI {
	return nil
}
