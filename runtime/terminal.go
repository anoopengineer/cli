package runtime

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Terminal is a representation of an actual command line terminal. Users should Initialize()
// before making any API calls. TearDown() should be called before the program exits
type Terminal interface {
	Initialize(prompt string) error
	SetPrompt(prompt string)

	PrintBuf([]byte) (int, error)
	Printf(format string, a ...interface{}) (int, error)
	Println(a ...interface{}) (int, error)

	ReadLine() (string, error)
	ReadPassword(prompt string) (password string, err error)
	ReadAnswer(question string) (answer string, err error)

	TearDown() error
}

// NewTerminal returns a default implementation of Terminal
func NewTerminal() Terminal {
	return &myTerminal{}
}

type myTerminal struct {
	oldState *terminal.State
	term     *terminal.Terminal
	prompt   string
}

func (m *myTerminal) Initialize(prompt string) error {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	m.oldState = oldState
	term := terminal.NewTerminal(os.Stdin, prompt)
	m.term = term
	m.prompt = prompt
	return nil
}

func (m *myTerminal) SetPrompt(prompt string) {
	m.term.SetPrompt(prompt)
	m.prompt = prompt
}

func (m *myTerminal) TearDown() error {
	return terminal.Restore(int(os.Stdin.Fd()), m.oldState)
}

func (m *myTerminal) ReadLine() (string, error) {
	if m.term == nil {
		return "", ErrTerminalNotInitialized
	}
	return m.term.ReadLine()
}

func (m *myTerminal) ReadPassword(prompt string) (string, error) {
	if m.term == nil {
		return "", ErrTerminalNotInitialized
	}
	return m.term.ReadPassword(prompt)
}

func (m *myTerminal) ReadAnswer(question string) (answer string, err error) {
	if m.term == nil {
		return "", ErrTerminalNotInitialized
	}
	oldPrompt := m.prompt
	m.term.SetPrompt(question)
	answer, err = m.ReadLine()
	m.term.SetPrompt(oldPrompt)
	return
}

func (m *myTerminal) PrintBuf(buf []byte) (int, error) {
	if m.term == nil {
		return -1, ErrTerminalNotInitialized
	}
	return m.term.Write(buf)
}

func (m *myTerminal) Printf(format string, a ...interface{}) (int, error) {
	if m.term == nil {
		return -1, ErrTerminalNotInitialized
	}
	buf := fmt.Sprintf(format, a...)
	return m.term.Write([]byte(buf))
}

func (m *myTerminal) Println(a ...interface{}) (int, error) {
	if m.term == nil {
		return -1, ErrTerminalNotInitialized
	}
	buf := fmt.Sprintln(a...)
	return m.term.Write([]byte(buf))
}
