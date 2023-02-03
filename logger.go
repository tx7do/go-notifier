package notifier

import "log"

type Logger interface {
	Info(fmt string, a ...interface{})
	Error(fmt string, a ...interface{})
}

type DefaultLogger struct {
}

func (l DefaultLogger) Info(fmt string, a ...interface{}) {
	log.Printf(fmt, a...)
}

func (l DefaultLogger) Error(fmt string, a ...interface{}) {
	log.Fatalf(fmt, a...)
}
