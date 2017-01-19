package runtime

import "errors"

var (
	// ErrTerminalNotInitialized is returned when any of the Terminal API is called without calling Initialize() first
	ErrTerminalNotInitialized = errors.New("Terminal is not initialized. Maybe you haven't called the Initialize() method first")
	// ErrUnknownCommand is returned when user enters a subcommand that is not implemented
	ErrUnknownCommand = errors.New("Unknown command. Type 'help' to see the list of available commands")
	// ErrNilConfig is returned when config passed to Run() is nil
	ErrNilConfig = errors.New("Config cannot be nil")
)
