package utils

import "testing"

func TestUtilityPropsResolve(t *testing.T) {
	tests := []struct {
		name  string
		props UtilityProps
		want  string
	}{
		{"empty", UtilityProps{}, ""},
		{"single field", UtilityProps{P: "4"}, "p-4"},
		{"multiple fields", UtilityProps{P: "4", Mx: "auto", Bg: "card"}, "p-4 mx-auto bg-card"},
		{"flex direction", UtilityProps{Flex: "col"}, "flex flex-col"},
		{"flex inline", UtilityProps{Flex: "inline"}, "inline-flex"},
		{"gap alias sm", UtilityProps{Gap: "sm"}, "gap-2"},
		{"gap alias lg", UtilityProps{Gap: "lg"}, "gap-6"},
		{"gap raw", UtilityProps{Gap: "3"}, "gap-3"},
		{"rounded default", UtilityProps{Rounded: "default"}, "rounded"},
		{"rounded lg", UtilityProps{Rounded: "lg"}, "rounded-lg"},
		{"shadow default", UtilityProps{Shadow: "true"}, "shadow"},
		{"border default", UtilityProps{Border: "true"}, "border"},
		{"border passthrough", UtilityProps{Border: "border-red-500"}, "border-red-500"},
		{"grow default", UtilityProps{Grow: "true"}, "grow"},
		{"shrink default", UtilityProps{Shrink: "default"}, "shrink"},
		{"hidden", UtilityProps{Hidden: true}, "hidden"},
		{"truncate", UtilityProps{Truncate: true}, "truncate"},
		{"grid + col", UtilityProps{Grid: "3", Col: "2"}, "grid-cols-3 col-span-2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.props.Resolve(); got != tt.want {
				t.Fatalf("Resolve() = %q, want %q", got, tt.want)
			}
		})
	}
}
