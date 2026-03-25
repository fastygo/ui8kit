package ui

import "github.com/fastygo/ui8kit/utils"

type BoxProps struct {
	utils.UtilityProps
	Class string
	Tag   string
}

type StackProps struct {
	utils.UtilityProps
	Class string
	Tag   string
}

type GroupProps struct {
	utils.UtilityProps
	Class string
	Tag   string
	Grow  bool
}

type ContainerProps struct {
	utils.UtilityProps
	Class string
}

type BlockProps = BoxProps

type ButtonProps struct {
	utils.UtilityProps
	Variant  string
	Size     string
	Href     string
	Class    string
	Type     string
	Disabled bool
}

type BadgeProps struct {
	utils.UtilityProps
	Variant string
	Size    string
	Class   string
}

type TextProps struct {
	utils.UtilityProps
	Class         string
	Tag           string
	FontSize      string
	FontWeight    string
	LineHeight    string
	LetterSpacing string
	TextColor     string
	TextAlign     string
	Truncate      bool
}

type TitleProps struct {
	utils.UtilityProps
	Class         string
	Order         int
	FontSize      string
	FontWeight    string
	LineHeight    string
	LetterSpacing string
	TextColor     string
	TextAlign     string
	Truncate      bool
}

type FieldOption struct {
	Value string
	Label string
}

type FieldProps struct {
	utils.UtilityProps
	Class        string
	Variant      string
	Size         string
	Type         string
	Name         string
	ID           string
	Placeholder  string
	Value        string
	Rows         int
	Min          string
	Max          string
	Checked      bool
	Disabled     bool
	Required     bool
	Autocomplete string
	Component    string
	Options      []FieldOption
}

type IconProps struct {
	Name  string
	Size  string
	Class string
}
