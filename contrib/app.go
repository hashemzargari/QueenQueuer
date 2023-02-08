package contrib

import (
	"github.com/hashemzargari/QueenQueuer/logging"
	loggingAdapters "github.com/hashemzargari/QueenQueuer/logging/adapters"
)

type OptionKind string

const (
	K_BootMode OptionKind = "kind_boot_mode"
	K_Broker   OptionKind = "kind_broker"
	K_Databse  OptionKind = "kind_database"
	K_Logger   OptionKind = "kind_logger"
)

type ConfigOption struct {
	Kind  OptionKind
	Value any
}

type App struct {
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

func (a *App) getDefaultConfigOptions() []*ConfigOption {
	var options []*ConfigOption
	// TODO: Add default options
	// Logger
	options = append(options, &ConfigOption{
		Kind:  K_Logger,
		Value: loggingAdapters.DefaultLogger{},
	})
	return options
}

func NewApp(options ...*ConfigOption) *App {
	a := &App{}
	allOptions := a.getDefaultConfigOptions()
	allOptions = append(allOptions, options...)
	a.updateConfig(allOptions...)
	return a
}
