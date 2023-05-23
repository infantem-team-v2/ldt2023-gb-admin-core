package tlogger

type LoggerMessenger interface {
	LoggerMessage() string
}

type ILogger interface {
	// sendLog Send log to remote logs storage
	sendLog(msg string, level LogLevel)

	// Default logger methods
	Infof(msgf string, args ...interface{})
	Debugf(msgf string, args ...interface{})
	Warnf(msgf string, args ...interface{})
	Errorf(msgf string, args ...interface{})
	// ErrorFull Logging error w/ stacktrace
	ErrorFull(err error, requestId string)
}
