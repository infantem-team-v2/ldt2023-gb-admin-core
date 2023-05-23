package _interface

type IError interface {
	LoggerMessage() string
	Wrap(err IError) IError
}
