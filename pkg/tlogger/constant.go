package tlogger

const (
	labels string = "source=%s job=error_logging statusCode=%d internalError=%s requestId=%s"
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)
