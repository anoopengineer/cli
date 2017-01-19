package runtime

import "github.com/gocommons/cli/color"

type Theme struct {
	PromptColor color.Color
	PromptStyle color.Style

	TraceColor color.Color
	TraceStyle color.Style

	InfoColor color.Color
	InfoStyle color.Style

	WarnColor color.Color
	WarnStyle color.Style

	ErrorColor color.Color
	ErrorStyle color.Style
}

func DefaultTheme() Theme {
	return Theme{
		PromptColor: color.Green,
		PromptStyle: color.Bold,

		TraceColor: color.LightGray,
		TraceStyle: color.Dim,

		InfoColor: color.DefaultForeground,
		InfoStyle: color.Normal,

		WarnColor: color.Magenta,
		WarnStyle: color.Normal,

		ErrorColor: color.Red,
		ErrorStyle: color.Bold,
	}
}
