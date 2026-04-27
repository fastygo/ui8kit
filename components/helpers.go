package components

import (
	"fmt"
	"strings"

	"github.com/fastygo/ui8kit/utils"
)

func cardClasses(p CardProps) string {
	return utils.Cn(utils.CardVariant(p.Variant), p.Class)
}

func cardTag(tag string) string {
	return utils.ResolveTag(tag, "div", utils.TagGroupLayout)
}

func cardTitleTag(order int) string {
	if order < 1 || order > 6 {
		return "h3"
	}
	return fmt.Sprintf("h%d", order)
}

func accordionType(value string) string {
	switch strings.TrimSpace(value) {
	case "multiple":
		return "multiple"
	default:
		return "single"
	}
}

func accordionState(open bool) string {
	if open {
		return "open"
	}
	return "closed"
}

func accordionTriggerID(value string) string {
	return "ui8kit-accordion-trigger-" + strings.TrimSpace(value)
}

func accordionPanelID(value string) string {
	return "ui8kit-accordion-panel-" + strings.TrimSpace(value)
}

func sheetID(value string) string {
	if strings.TrimSpace(value) == "" {
		return "sheet"
	}
	return strings.TrimSpace(value)
}

func sheetPanelID(value string) string {
	return sheetID(value) + "-panel"
}

func sheetTitleID(value string) string {
	return sheetID(value) + "-title"
}

func sheetContentID(value string) string {
	return sheetID(value) + "-content"
}

func sheetPanelClasses(p SheetProps) string {
	return utils.Cn(
		"ui-sheet-panel",
		utils.SheetSideVariant(p.Side),
		utils.SheetSizeVariant(p.Size),
		p.Class,
	)
}

func dialogID(value string) string {
	id := strings.TrimSpace(value)
	if id == "" {
		return "ui8kit-dialog"
	}
	return "ui8kit-dialog-" + id
}

func dialogTitleID(dialogIDValue string) string {
	return dialogID(dialogIDValue) + "-title"
}

func dialogDescriptionID(dialogIDValue string) string {
	return dialogID(dialogIDValue) + "-description"
}

func alertVariantClass(value string) string {
	switch strings.TrimSpace(value) {
	case "destructive":
		return "ui-alert-destructive"
	case "success":
		return "ui-alert-success"
	case "warning":
		return "ui-alert-warning"
	default:
		return "ui-alert-default"
	}
}

func tabsID(value string) string {
	if strings.TrimSpace(value) == "" {
		return "ui8kit-tabs"
	}
	return "ui8kit-tabs-" + strings.TrimSpace(value)
}

func tabButtonID(tabs string, value string) string {
	id := strings.TrimSpace(value)
	if id == "" {
		id = "tab"
	}
	return tabsID(tabs) + "-trigger-" + id
}

func tabPanelID(tabs string, value string) string {
	id := strings.TrimSpace(value)
	if id == "" {
		id = "tab"
	}
	return tabsID(tabs) + "-panel-" + id
}

func comboboxID(value string) string {
	id := strings.TrimSpace(value)
	if id == "" {
		return "ui8kit-combobox"
	}
	return "ui8kit-combobox-" + id
}

func comboboxInputID(value string) string {
	return comboboxID(value) + "-input"
}

func comboboxOptionsID(value string) string {
	return comboboxID(value) + "-options"
}

func comboboxOptionID(root string, option string) string {
	return comboboxID(root) + "-option-" + strings.TrimSpace(option)
}

func tooltipID(value string) string {
	id := strings.TrimSpace(value)
	if id == "" {
		return "ui8kit-tooltip"
	}
	return "ui8kit-tooltip-" + id
}

func tooltipContentID(value string) string {
	return tooltipID(value) + "-content"
}
