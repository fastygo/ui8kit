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

type BlockProps struct {
	ID    string
	Class string
	Tag   string
}

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

type PictureProps struct {
	Class string
}

type SourceProps struct {
	SrcSet string
	Src    string
	Media  string
	Type   string
	Sizes  string
}

type FigureProps struct {
	Class string
}

type FigureCaptionProps struct {
	Class string
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

type TableProps struct {
	Class string
	Attrs templ.Attributes
}

type TableCaptionProps struct {
	Class string
}

type TableSectionProps struct {
	Class string
}

type TableRowProps struct {
	Class string
}

type TableCellProps struct {
	Class   string
	Scope   string
	ColSpan int
	RowSpan int
	Headers string
	Abbr    string
}

type TableColGroupProps struct {
	Class string
	Span  int
}

type TableColProps struct {
	Class string
	Span  int
}

type ListProps struct {
	Class string
	Tag   string
}

type ListItemProps struct {
	Class string
	Value int
}

type DescriptionListProps struct {
	Class string
}

type DescriptionTermProps struct {
	Class string
}

type DescriptionDetailsProps struct {
	Class string
}

type DisclosureProps struct {
	Class string
	Open  bool
}

type DisclosureSummaryProps struct {
	Class string
}

type FormProps struct {
	ID           string
	Class        string
	Action       string
	Method       string
	Enctype      string
	Autocomplete string
	Name         string
	Target       string
	NoValidate   bool
	Attrs        templ.Attributes
}

type FieldsetProps struct {
	Class    string
	Name     string
	Form     string
	Disabled bool
	Attrs    templ.Attributes
}

type LegendProps struct {
	Class string
}

type DataListProps struct {
	ID    string
	Class string
}

type OptGroupProps struct {
	Class    string
	Label    string
	Disabled bool
}

type OutputProps struct {
	ID    string
	Class string
	Name  string
	For   string
	Value string
	Attrs templ.Attributes
}

type MeterProps struct {
	ID      string
	Class   string
	Value   string
	Min     string
	Max     string
	Low     string
	High    string
	Optimum string
	Attrs   templ.Attributes
}

type ProgressProps struct {
	ID    string
	Class string
	Value string
	Max   string
	Attrs templ.Attributes
}
