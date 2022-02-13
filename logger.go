package main

import (
	"log"
	"os"
)

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type logger struct {
	lg *log.Logger
}

func NewLogger() logger {
	return logger{
		lg: log.New(os.Stdout, "", 5),
	}
}

func (l logger) Log(args ...interface{}) {
	l.lg.Println(args...)
}

func (l logger) Logf(format string, args ...interface{}) {
	l.lg.Printf(format, args...)
}
