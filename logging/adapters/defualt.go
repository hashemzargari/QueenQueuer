package adapters

import (
	"github.com/fatih/color"
	"github.com/hashemzargari/QueenQueuer/logging"
	"log"
	"runtime"
	"time"
)

type DefaultLogger struct {
}

func (d DefaultLogger) bindCallers() []*logging.Caller {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	var callers []*logging.Caller
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		callers = append(callers, &logging.Caller{
			Function: f.Name(),
			File:     file,
			Line:     line,
		})
	}
	return callers
}

func (d DefaultLogger) getStackTrace() string {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, true)
	return string(buf[:n])
}

func (d DefaultLogger) Error(err error) {
	d.prepareLog(err.Error(), logging.ERROR)
}

func (d DefaultLogger) Message(msg string) {
	d.prepareLog(msg, logging.INFO)
}
func (d DefaultLogger) Warning(msg string) {
	d.prepareLog(msg, logging.WARNING)
}

func (d DefaultLogger) prepareLog(msg string, logMode logging.LogMode) {
	callers := d.bindCallers()
	d.saveLog(&logging.Log{
		Time:      time.Now(),
		Message:   msg,
		Callers:   callers,
		TraceBack: d.getStackTrace(),
	}, logMode)
}

func (d DefaultLogger) saveLog(logPointer *logging.Log, logMode logging.LogMode) {
	// TODO: add handlers and routing logic for logs
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)
	cl := color.New(color.FgCyan).SprintFunc()
	switch logMode {
	case logging.WARNING:
		cl = color.New(color.FgYellow).SprintFunc()
	case logging.ERROR:
		cl = color.New(color.FgRed).SprintFunc()
	}
	log.Println(cl("[", logMode, "]\t", logPointer.Message))
	log.Println(cl("Full stack trace:\n", logPointer.TraceBack))
}
