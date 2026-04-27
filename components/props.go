package components

type CardProps struct {
	Class   string
	Variant string
	Tag     string
}

type CardHeaderProps struct {
	Class string
}

type CardTitleProps struct {
	Class string
	Order int
}

type CardDescriptionProps struct {
	Class string
}

type CardContentProps struct {
	Class string
}

type CardFooterProps struct {
	Class string
}

type AccordionProps struct {
	Class string
	Type  string
}

type AccordionItemProps struct {
	Class string
	Value string
	Open  bool
}

type AccordionTriggerProps struct {
	Class string
	Value string
	Open  bool
}

type AccordionContentProps struct {
	Class string
	Value string
	Open  bool
}

type SheetProps struct {
	Class string
	ID    string
	Side  string
	Size  string
	Title string
}
