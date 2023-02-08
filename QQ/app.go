package QQ

import (
	"github.com/hashemzargari/QueenQueuer/logging"
	loggingAdapters "github.com/hashemzargari/QueenQueuer/logging/adapters"
)

type OptionKind string

const (
	K_Broker  OptionKind = "broker"
	K_Databse OptionKind = "database"
	K_Logger  OptionKind = "kind_logger"
)

type ConfigOption struct {
	Kind  OptionKind
	Value any
}

func LoggerConfig(logger logging.Logger) *ConfigOption {
	return &ConfigOption{
		Kind:  K_Logger,
		Value: logger,
	}
}

type App struct {
	Name   string
	logger logging.Logger
	tasks  map[string]Task
}

func (a *App) SetConfig(option *ConfigOption) *App {
	a.updateConfig(option)
	return a
}

func (a *App) Logger() logging.Logger {
	return a.logger
}

func (a *App) RegisterTasks(tasks ...Task) {
	//TODO implement me
	panic("implement me")
}

func (a *App) updateConfig(options ...*ConfigOption) {
	for _, option := range options {
		switch option.Kind {
		case K_Logger:
			a.logger = option.Value.(logging.Logger)
		}
	}
}

func (a *App) prepareDefaultConfigOptions() *App {
	// TODO: Add default options
	// Logger
	a.
		SetConfig(LoggerConfig(loggingAdapters.DefaultLogger{}))
	return a
}

func NewApp(name string, options ...*ConfigOption) *App {
	a := &App{
		Name: name,
	}
	a.
		prepareDefaultConfigOptions().
		updateConfig(options...)
	return a
}
