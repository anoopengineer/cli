package commands

import (
	"fmt"
	"strings"

	"github.com/gocommons/cli/runtime"
)

// Quit is a default command that gets added to all application.
// It allows the user to exit the prompt.
type Quit struct {
}

// Name is the subcommand name
func (c *Quit) Name() string {
	return "quit"
}

// Run gets executed when user runs "$commandName quit"
func (c *Quit) Run(runtime runtime.Runtime, args []string) (exitProgram bool, err error) {
	fmt.Println()
	return true, nil
}

// Synopsis is a short help string
func (c *Quit) Synopsis() string {
	return "quits the program"
}

// Help shows a short help message
func (c *Quit) Help() string {
	helpText := `quits the program.

`
	return strings.TrimSpace(helpText)
}
