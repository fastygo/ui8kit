package ui8kit

import (
	"regexp"
	"testing"
)

func TestVersionFormat(t *testing.T) {
	semver := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !semver.MatchString(Version) {
		t.Fatalf("Version %q does not match semver format (expected X.Y.Z)", Version)
	}
}

func TestVersionNotEmpty(t *testing.T) {
	if Version == "" {
		t.Fatal("Version must not be empty")
	}
}
