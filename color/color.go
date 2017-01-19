package color

import "fmt"

// Color represents a Foreground or Background color value.
// Refer to http://misc.flogisoft.com/bash/tip_colors_and_formatting to know more about color in terminals
type Color int

// Style represents attributes like bold, italics etc that can be associated to text in terminal
type Style int

var (
	// NoColor represents NoColor
	NoColor Color = -1
	// DefaultForeground represents DefaultForeground
	DefaultForeground Color = 39
	// Black represents Black
	Black Color = 30
	// Red represents Red
	Red Color = 31
	// Green represents Green
	Green Color = 32
	// Yellow represents Yellow
	Yellow Color = 33
	// Blue represents Blue
	Blue Color = 34
	// Magenta represents Magenta
	Magenta Color = 35
	// Cyan represents Cyan
	Cyan Color = 36
	// LightGray represents LightGray
	LightGray Color = 37
	// DarkGray represents DarkGray
	DarkGray Color = 90
	// LightRed represents LightRed
	LightRed Color = 91
	// LightGreen represents LightGreen
	LightGreen Color = 92
	// LightYellow represents LightYellow
	LightYellow Color = 93
	// LightBlue represents LightBlue
	LightBlue Color = 94
	// LightMagenta represents LightMagenta
	LightMagenta Color = 95
	// LightCyan represents LightCyan
	LightCyan Color = 96
	// White represents White
	White Color = 97
)

var (
	// Normal represents Normal
	Normal Style
	// Bold represents Bold
	Bold Style = 1
	// Dim represents Dim
	Dim Style = 2
	// Italics represents Italics
	Italics Style = 3
	// Underline represents Underline
	Underline Style = 4
	// Blink represents Blink
	Blink Style = 5
	// Inverted represents Inverted
	Inverted Style = 7
	// Hidden represents Hidden
	Hidden Style = 8
)

// InColor converts into a plain string into a string that gets rendered with color in a terminal
func InColor(str string, color Color, style Style) string {
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", style, color, str)
}
