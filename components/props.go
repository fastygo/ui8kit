package components

import "github.com/fastygo/ui8kit/utils"

type CardProps struct {
	utils.UtilityProps
	Class   string
	Variant string
	Tag     string
}

type CardHeaderProps struct {
	utils.UtilityProps
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
	utils.UtilityProps
	Class string
}

type CardFooterProps struct {
	utils.UtilityProps
	Class string
}

type AccordionProps struct {
	utils.UtilityProps
	Class string
	Type  string
}

type AccordionItemProps struct {
	utils.UtilityProps
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
	utils.UtilityProps
	Class string
	Value string
	Open  bool
}

type SheetProps struct {
	utils.UtilityProps
	Class string
	ID    string
	Side  string
	Size  string
	Title string
}
