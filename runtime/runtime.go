package runtime

import (
	"fmt"

	"github.com/gocommons/cli/color"
)

// Runtime provides a set of helper APIs that makes the life easy for a developer.
//
// Trace(), Info(), Warn() and Error() provide out of the box color coded debug statements
// to the console.
//
// Term() object provides even low level APIs like ReadLine(), ReadPassword() and ReadAnswer()
//
// OSArgument() provides you the entire os.Args fields that was received
//
//CommandArgs is the part without the sub command name and sub command arguments.
// For example, if the program was invoked as
//	$mycli -a arg1 -b arg2 myCommand -x commangArg1 -y commandArg2 -z commandArg3
// CommandArgs() will return []{"-a", "arg1", "-b", "arg2"}
//
// GetCommand returns a Command object and is useful if you want to invoke another subcommand
// from within a subcommand's Run()
type Runtime interface {
	Term() Terminal
	OSArgument() []string
	CommandArgs() []string
	GetAllCommandNames() []string
	GetCommand(name string) Command
	RunCommand(name string, args []string) (exitProgram bool, err error)
	Config() Config
	Store() map[string]interface{}

	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
}

// GetDefaultRuntime provides a default implementation of Runtime interface
func GetDefaultRuntime(term Terminal, osArguments []string, commandArguments []string, config Config, store map[string]interface{}) Runtime {
	return &defaultRuntime{
		term:             term,
		oSArguments:      osArguments,
		commandArguments: commandArguments,
		config:           config,
		store:            store,
	}
}

type defaultRuntime struct {
	term             Terminal
	oSArguments      []string
	commandArguments []string
	config           Config
	store            map[string]interface{}
}

func (c *defaultRuntime) Term() Terminal {
	return c.term
}

func (c *defaultRuntime) OSArgument() []string {
	return c.oSArguments
}

func (c *defaultRuntime) CommandArgs() []string {
	return c.commandArguments
}

func (c *defaultRuntime) Config() Config {
	return c.config
}

func (c *defaultRuntime) Store() map[string]interface{} {
	return c.store
}

func (c *defaultRuntime) GetCommand(name string) Command {
	return commandMap[name]
}

func (c *defaultRuntime) GetAllCommandNames() []string {
	keys := make([]string, 0, len(commandMap))
	for k := range commandMap {
		keys = append(keys, k)
	}
	return keys
}

func (c *defaultRuntime) RunCommand(name string, args []string) (exitProgram bool, err error) {
	cmd, ok := commandMap[name]
	if !ok {
		err = ErrUnknownCommand
		exitProgram = false
		return
	}
	return cmd.Run(c, args)
}

func (c *defaultRuntime) Trace(format string, a ...interface{}) {
	txt := fmt.Sprintf(format, a...)
	c.term.Println(color.InColor(txt, c.config.Theme().TraceColor, c.config.Theme().TraceStyle))
}

func (c *defaultRuntime) Info(format string, a ...interface{}) {
	txt := fmt.Sprintf(format, a...)
	c.term.Println(color.InColor(txt, c.config.Theme().InfoColor, c.config.Theme().InfoStyle))
}

func (c *defaultRuntime) Warn(format string, a ...interface{}) {
	txt := fmt.Sprintf(format, a...)
	c.term.Println(color.InColor(txt, c.config.Theme().WarnColor, c.config.Theme().WarnStyle))
}

func (c *defaultRuntime) Error(format string, a ...interface{}) {
	txt := fmt.Sprintf(format, a...)
	c.term.Println(color.InColor(txt, c.config.Theme().ErrorColor, c.config.Theme().ErrorStyle))
}
