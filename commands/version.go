package commands

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gocommons/cli/runtime"
)

// Version is a default command that gets added to all application.
// It displays the version information of the application.
type Version struct {
}

// Name is the subcommand name
func (c *Version) Name() string {
	return "version"
}

// Run gets executed when user runs "$commandName version"
func (c *Version) Run(runtime runtime.Runtime, args []string) (exitProgram bool, err error) {
	var buf bytes.Buffer
	appName := runtime.Config().AppName()
	versionNumber := runtime.Config().VersionNumber()
	versionTag := runtime.Config().VersionTag()
	buf.WriteString(fmt.Sprintf("%s - %s", appName, versionNumber))
	if versionTag != "" {
		buf.WriteString(fmt.Sprintf(" (%s)", versionTag))
	}
	runtime.Info(buf.String())
	return false, nil
}

// Synopsis is a short help string
func (c *Version) Synopsis() string {
	return "Displays version information"
}

// Help shows a short help message
func (c *Version) Help() string {
	helpText := `Displays version information.

`
	return strings.TrimSpace(helpText)
}
