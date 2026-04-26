package ui

import (
	"github.com/a-h/templ"
	c "github.com/fastygo/ui8kit/components"
)

func Card(props CardProps) templ.Component {
	return c.Card(cardPropsFromUI(props))
}

func CardHeader(props CardHeaderProps) templ.Component {
	return c.CardHeader(cardHeaderPropsFromUI(props))
}

func CardTitle(props CardTitleProps, value string) templ.Component {
	return c.CardTitle(cardTitlePropsFromUI(props), value)
}

func CardDescription(props CardDescriptionProps, value string) templ.Component {
	return c.CardDescription(cardDescriptionPropsFromUI(props), value)
}

func CardContent(props CardContentProps) templ.Component {
	return c.CardContent(cardContentPropsFromUI(props))
}

func CardFooter(props CardFooterProps) templ.Component {
	return c.CardFooter(cardFooterPropsFromUI(props))
}

func Accordion(props AccordionProps) templ.Component {
	return c.Accordion(accordionPropsFromUI(props))
}

func AccordionItem(props AccordionItemProps) templ.Component {
	return c.AccordionItem(accordionItemPropsFromUI(props))
}

func AccordionTrigger(props AccordionTriggerProps) templ.Component {
	return c.AccordionTrigger(accordionTriggerPropsFromUI(props))
}

func AccordionContent(props AccordionContentProps) templ.Component {
	return c.AccordionContent(accordionContentPropsFromUI(props))
}

func Sheet(props SheetProps) templ.Component {
	return c.Sheet(sheetPropsFromUI(props))
}

func cardPropsFromUI(props CardProps) c.CardProps {
	return c.CardProps{
		Class:   props.Class,
		Variant: props.Variant,
		Tag:     props.Tag,
	}
}

func cardHeaderPropsFromUI(props CardHeaderProps) c.CardHeaderProps {
	return c.CardHeaderProps{
		Class: props.Class,
	}
}

func cardTitlePropsFromUI(props CardTitleProps) c.CardTitleProps {
	return c.CardTitleProps{
		Class: props.Class,
		Order: props.Order,
	}
}

func cardDescriptionPropsFromUI(props CardDescriptionProps) c.CardDescriptionProps {
	return c.CardDescriptionProps{
		Class: props.Class,
	}
}

func cardContentPropsFromUI(props CardContentProps) c.CardContentProps {
	return c.CardContentProps{
		Class: props.Class,
	}
}

func cardFooterPropsFromUI(props CardFooterProps) c.CardFooterProps {
	return c.CardFooterProps{
		Class: props.Class,
	}
}

func accordionPropsFromUI(props AccordionProps) c.AccordionProps {
	return c.AccordionProps{
		Class: props.Class,
		Type:  props.Type,
	}
}

func accordionItemPropsFromUI(props AccordionItemProps) c.AccordionItemProps {
	return c.AccordionItemProps{
		Class: props.Class,
		Value: props.Value,
		Open:  props.Open,
	}
}

func accordionTriggerPropsFromUI(props AccordionTriggerProps) c.AccordionTriggerProps {
	return c.AccordionTriggerProps{
		Class: props.Class,
		Value: props.Value,
		Open:  props.Open,
	}
}

func accordionContentPropsFromUI(props AccordionContentProps) c.AccordionContentProps {
	return c.AccordionContentProps{
		Class: props.Class,
		Value: props.Value,
		Open:  props.Open,
	}
}

func sheetPropsFromUI(props SheetProps) c.SheetProps {
	return c.SheetProps{
		Class: props.Class,
		ID:    props.ID,
		Side:  props.Side,
		Size:  props.Size,
		Title: props.Title,
	}
}
