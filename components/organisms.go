package components

import "github.com/fastygo/ui8kit/utils"

type DialogProps struct {
	utils.UtilityProps
	ID    string
	Class string
	Title string
	Open  bool
}

type DialogTriggerProps struct {
	For   string
	Label string
	Class string
}

type AlertDialogProps struct {
	utils.UtilityProps
	ID    string
	Class string
	Open  bool
}

type AlertProps struct {
	utils.UtilityProps
	Class   string
	Variant string
}

type BreadcrumbProps struct {
	utils.UtilityProps
	Class string
	Items []BreadcrumbItem
}

type BreadcrumbItem struct {
	Label    string
	Href     string
	Current  bool
	Disabled bool
}

type TabsProps struct {
	utils.UtilityProps
	ID        string
	Class     string
	Tabs      []TabDescriptor
	ActiveTab string
}

type TabDescriptor struct {
	Value    string
	Label    string
	Disabled bool
}

type TabsPanelProps struct {
	TabsID string
	Value  string
	Class  string
	Active bool
}

type ComboboxProps struct {
	utils.UtilityProps
	ID          string
	Class       string
	Name        string
	Label       string
	Placeholder string
	Value       string
	Options     []ComboboxOptionProps
}

type ComboboxOptionProps struct {
	Value    string
	Label    string
	Disabled bool
}

type TooltipProps struct {
	utils.UtilityProps
	ID    string
	Class string
	Text  string
	Open  bool
}
