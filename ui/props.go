package ui

import "github.com/a-h/templ"

type DOMProps struct {
	ID       string
	Role     string
	TabIndex string
	Attrs    templ.Attributes
}

type BoxProps struct {
	Class string
	Tag   string
}

type StackProps struct {
	Class string
	Tag   string
}

type GroupProps struct {
	Class string
	Tag   string
	Grow  bool
}

type ContainerProps struct {
	Class string
	Tag   string
}

type BlockProps = BoxProps

type ButtonProps struct {
	ID        string
	Role      string
	TabIndex  string
	Attrs     templ.Attributes
	Variant   string
	Size      string
	Href      string
	Class     string
	Type      string
	AriaLabel string
	Disabled  bool
}

type BadgeProps struct {
	Variant string
	Size    string
	Class   string
}

type TextProps struct {
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
	ID           string
	Role         string
	TabIndex     string
	Attrs        templ.Attributes
	Class        string
	Variant      string
	Size         string
	Type         string
	Name         string
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
	Label        string
	Hint         string
	Error        string
	AriaLabel    string
	Switch       bool
}

type IconProps struct {
	Name  string
	Size  string
	Class string
}

type ImageProps struct {
	Class    string
	Src      string
	Alt      string
	Width    string
	Height   string
	Fit      string
	Position string
	Aspect   string
	Loading  string
}

type GridProps struct {
	Class string
	Cols  string
}

type GridColProps struct {
	Class string
	Span  int
	Start int
	End   int
	Order int
}
