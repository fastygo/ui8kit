package components

type DialogProps struct {
	ID          string
	Class       string
	Title       string
	Description string
	AriaLabel   string
	Open        bool
}

type DialogTriggerProps struct {
	For   string
	Label string
	Class string
	Open  bool
}

type AlertDialogProps struct {
	ID        string
	Class     string
	AriaLabel string
	Open      bool
}

type AlertProps struct {
	Class   string
	Variant string
}

type BreadcrumbProps struct {
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
	ID    string
	Class string
	Text  string
	Open  bool
}
