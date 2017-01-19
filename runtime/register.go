package runtime

import "sync"

var commandMap map[string]Command

var once sync.Once

// RegisterCommand adds the command to the list of invokable subcommands.
// Your subcommand is available to the user only if you have Registered the
// command using this API.
//
// Best place to register would be to define an
// init function in your command struct and call this API from it.
//
// 	func init() {
//		cli.RegisterCommand(&MySubCommand{})
//	}
func RegisterCommand(cmd Command) {
	once.Do(func() {
		commandMap = make(map[string]Command)
	})
	commandMap[cmd.Name()] = cmd
}

func GetCommandMap() map[string]Command {
	return commandMap
}
