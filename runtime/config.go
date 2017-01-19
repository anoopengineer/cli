package runtime

import (
	"github.com/gocommons/cli/color"
	"github.com/gocommons/cli/util"
)

// Config represents the basic configuration parameters that change the runtime
// behavior of the application.
type Config interface {
	AppName() string
	SetPrompt(prompt string)
	Prompt() string
	VersionNumber() string
	VersionTag() string
	Get(key string) (val interface{}, ok bool)
	Set(key string, val interface{})
	SetTheme(theme Theme)
	Theme() Theme
}

// GetDefaultConfig returns an implementation of Config interface where the prompt is
// set to the binary name.
func GetDefaultConfig(appName, versionNumber, versionTag string) Config {
	theme := DefaultTheme()
	return &defaultConfig{
		appName:       appName,
		prompt:        color.InColor(util.GetBinaryName()+"> ", theme.PromptColor, theme.PromptStyle), //setting binary name as the prompt
		versionNumber: versionNumber,
		versionTag:    versionTag,
		bag:           map[string]interface{}{},
		theme:         theme,
	}
}

type defaultConfig struct {
	appName       string
	prompt        string
	versionNumber string
	versionTag    string
	bag           map[string]interface{}
	theme         Theme
}

func (c *defaultConfig) AppName() string {
	if c.appName == "" {
		c.appName = util.GetBinaryName()
	}
	return c.appName
}

func (c *defaultConfig) SetPrompt(prompt string) {
	c.prompt = prompt
}

func (c *defaultConfig) Prompt() string {
	return c.prompt
}

func (c *defaultConfig) VersionNumber() string {
	return c.versionNumber
}

func (c *defaultConfig) VersionTag() string {
	return c.versionTag
}

func (c *defaultConfig) Get(key string) (interface{}, bool) {
	val, ok := c.bag[key]
	return val, ok
}

func (c *defaultConfig) Set(key string, val interface{}) {
	c.bag[key] = val
}

func (c *defaultConfig) SetTheme(theme Theme) {
	c.theme = theme
}
func (c *defaultConfig) Theme() Theme {
	return c.theme
}
