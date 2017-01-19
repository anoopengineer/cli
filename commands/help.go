package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gocommons/cli/runtime"
	"github.com/gocommons/cli/util"
)

// Help is a default command that gets added to all application.
// It shows help details for all sub commands
type Help struct {
}

// Name is the subcommand name
func (c *Help) Name() string {
	return "help"
}

// Run gets executed when user runs "$commandName Help"
func (c *Help) Run(runtime runtime.Runtime, args []string) (exitProgram bool, err error) {
	if len(args) == 0 {
		showDefaultHelp(runtime)
		return false, nil
	}
	return showHelp(runtime, args)
}

func showDefaultHelp(r runtime.Runtime) {
	keys := r.GetAllCommandNames()
	sort.Strings(keys)
	maxLen := 0
	for _, v := range keys {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	//%-15s - %s
	formatString := fmt.Sprintf("  %%-%ds - %%s", maxLen)
	binName := util.GetBinaryName()
	r.Info("Usage: %s <command arguments> [subcommand name]  <subcommand arguments>", binName)
	r.Info("You can omit the subcommand name to get into %s shell", binName)
	if len(keys) > 0 {
		r.Info("Available sub-commands are:")
	}

	for _, v := range keys {
		c := r.GetCommand(v)
		r.Info(formatString, c.Name(), c.Synopsis())
	}
}

func showHelp(r runtime.Runtime, args []string) (exitProgram bool, err error) {
	if len(args) <= 0 {
		return false, nil
	}
	arg := args[0]
	cmd := r.GetCommand(arg)
	if cmd == nil {
		r.Error("Command not found %s", arg)
		return false, nil
	}
	r.Info(cmd.Help())

	return false, nil
}

// Synopsis is a short help string
func (c *Help) Synopsis() string {
	return "Prints the list of available subcommands and their synopsis. Use help <subcommand name> to know more about that subcommand"
}

// Help shows a short help message
func (c *Help) Help() string {
	helpText := `Helps the program.

`
	return strings.TrimSpace(helpText)
}
