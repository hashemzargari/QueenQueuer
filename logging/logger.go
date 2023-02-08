package logging

import (
	"time"
)

type LogMode string

const (
	ERROR   LogMode = "error"
	WARNING LogMode = "warning"
	INFO    LogMode = "info"
)

type Logger interface {
	Message(msg string)
	Warning(msg string)
	Error(err error)
}

type Caller struct {
	Function string
	File     string
	Line     int
}

type Log struct {
	Time      time.Time
	Message   string
	Callers   []*Caller
	TraceBack string
	LogMode   LogMode
}

// TODO: add routing and handler structs
