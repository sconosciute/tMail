package styles

import "github.com/charmbracelet/lipgloss"

var CurrentTheme = DefaultTheme

type Theme struct {
	Color  *Colors
	Styles map[StyleName]lipgloss.Style
}

type Colors struct {
	Primary   lipgloss.AdaptiveColor
	Secondary lipgloss.AdaptiveColor
	Accent    lipgloss.AdaptiveColor
}

//region styleNames

type StyleName int

const (
	TextInput StyleName = iota
	Paragraph
)

//endregion

//region defaultStyle

var DefaultTheme Theme

var (
	PrimaryColor = lipgloss.AdaptiveColor{
		Light: "#343434",
		Dark:  "#FEFCDB"}
	SecondaryColor = lipgloss.AdaptiveColor{
		Light: "2D4059",
		Dark:  "F7931E",
	}
	AccentColor = lipgloss.AdaptiveColor{
		Light: "EA5455",
		Dark:  "F05A28",
	}

	GlobalPad  = 1
	FieldWidth = 80
)

func InitDefaultTheme() {
	var defaultStyles = map[StyleName]lipgloss.Style{
		TextInput: lipgloss.NewStyle().
			BorderForeground(AccentColor).
			BorderStyle(lipgloss.NormalBorder()).
			Padding(GlobalPad).
			Width(FieldWidth),
	}

	DefaultTheme = Theme{
		Color: &Colors{
			Primary:   PrimaryColor,
			Secondary: SecondaryColor,
			Accent:    AccentColor,
		},
	}
	DefaultTheme.Styles = defaultStyles
}

var ()

//endregion
