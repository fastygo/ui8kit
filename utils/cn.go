package utils

import "strings"

// Cn joins non-empty class fragments with a single space.
func Cn(classes ...string) string {
	parts := make([]string, 0, len(classes))
	for _, c := range classes {
		if t := strings.TrimSpace(c); t != "" {
			parts = append(parts, t)
		}
	}
	return strings.Join(parts, " ")
}
