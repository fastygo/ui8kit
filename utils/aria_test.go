package utils

import "testing"

func TestAriaHelpers(t *testing.T) {
	if AriaExpanded(true)["aria-expanded"] != "true" {
		t.Fatal("AriaExpanded(true) should set aria-expanded=true")
	}
	if AriaExpanded(false)["aria-expanded"] != "false" {
		t.Fatal("AriaExpanded(false) should set aria-expanded=false")
	}
	if AriaControls("panel-1")["aria-controls"] != "panel-1" {
		t.Fatal("AriaControls should set aria-controls")
	}
	if _, ok := AriaControls("")["aria-controls"]; ok {
		t.Fatal("AriaControls should skip empty id")
	}
	if AriaLabelledBy("t")["aria-labelledby"] != "t" {
		t.Fatal("AriaLabelledBy should set aria-labelledby")
	}
	if AriaDescribedBy("d")["aria-describedby"] != "d" {
		t.Fatal("AriaDescribedBy should set aria-describedby")
	}
	if AriaModal(true)["aria-modal"] != "true" {
		t.Fatal("AriaModal should set aria-modal")
	}
	if AriaHidden(true)["aria-hidden"] != "true" {
		t.Fatal("AriaHidden should set aria-hidden")
	}
	if AriaSelected(false)["aria-selected"] != "false" {
		t.Fatal("AriaSelected should set aria-selected")
	}
	if AriaDisabled(true)["aria-disabled"] != "true" {
		t.Fatal("AriaDisabled should set aria-disabled")
	}
	if AriaRequired(true)["aria-required"] != "true" {
		t.Fatal("AriaRequired should set aria-required")
	}
	if AriaLive("assertive")["aria-live"] != "assertive" {
		t.Fatal("AriaLive(assertive) should set assertive")
	}
	if AriaLive("unknown")["aria-live"] != "polite" {
		t.Fatal("AriaLive(unknown) should fallback to polite")
	}
}
