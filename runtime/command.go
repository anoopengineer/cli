package runtime

// Command represents a runnable subcommand
// e.g. if you are implementing a git command line tool and you intend to
// invoke it as "git init", "git clone <args>" etc, then init and clone represents
// two Commands for your program
//
// To create a new sub command, implement this interface and have an init() method
// that calls
//	cli.RegisterCommand(&MyCommand{})
type Command interface {
	Name() string
	Run(runtime Runtime, args []string) (exitProgram bool, err error)
	Synopsis() string
	Help() string
}
