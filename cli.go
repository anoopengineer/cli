package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gocommons/cli/commands"
	"github.com/gocommons/cli/runtime"
)

/*
Run is the main entry point for your command line app.
Normally you would pass os.Args as args parameter to Run().
config object can be obtained using cli.GetDefaultConfig()
The following is a sample syntax that the command line parser understands

 $mycli -a arg1 -b arg2 myCommand -x commangArg1 -y commandArg2 -z commandArg3

Points to note:

• Commands can have arguments and also subcommands

• In the above example -a and -b are command arguments

• myCommand is a subcommand

• -x, -y and -z are arguments to the subcommand

The order in which the arguments are provided in the cli is important. It should always be command name, command arguments, subcommand name and subcommand arguments.
Anything that comes after the subcommand name is provided as arguments to the subcommand's Run() method.

*/
func Run(args []string, conf runtime.Config) error {
	if conf == nil {
		return runtime.ErrNilConfig
	}
	term := runtime.NewTerminal()
	if err := term.Initialize(conf.Prompt()); err != nil {
		panic(err)
	}
	defer term.TearDown()

	initializeDefaultCommands([]runtime.Command{&commands.Version{}, &commands.Quit{}, &commands.Help{}})

	commandArgs, subCommand, subcommandArgs := parseCommands(os.Args)
	runtime := runtime.GetDefaultRuntime(term, os.Args, commandArgs, conf, map[string]interface{}{})

	if subCommand != "" {
		//a subcommand is mentioned in the command prompt. Execute that command and exit.
		_, err := runtime.RunCommand(subCommand, subcommandArgs)
		if err != nil {
			runtime.Error(err.Error())
		}
		return err
	}
	runtime.Info("Type 'help' to see all available commands. Type 'quit' or Ctrl + d to exit")

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			fmt.Println()
			return io.EOF
		}
		if err != nil {
			runtime.Error("Error occurred", err)
			return err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		subCommandName := parts[0]
		subCommandArgs := parts[1:]

		exitProgram, err := runtime.RunCommand(subCommandName, subCommandArgs)
		if exitProgram {
			if err != nil {
				runtime.Error("Error occured: ", err.Error())
			}
			return err
		} else if err != nil {
			runtime.Error(err.Error())
		}
	}
}

func parseCommands(args []string) (commandArgs []string, subCommand string, subCommandArgs []string) {
	commandMap := runtime.GetCommandMap()
	for i, v := range args[1:] { //first element will be the binary name
		if _, ok := commandMap[v]; ok {
			return args[1 : i+1], v, args[i+2:]
		}
	}
	return nil, "", nil
}

func initializeDefaultCommands(commands []runtime.Command) {
	for _, v := range commands {
		runtime.RegisterCommand(v)
	}
}
