package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/cmd/set-version <version>")
		os.Exit(1)
	}

	version := os.Args[1]
	src, err := os.ReadFile("ui8kit.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read ui8kit.go: %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`const Version = "\d+\.\d+\.\d+"`)
	replacement := fmt.Sprintf(`const Version = "%s"`, version)
	updated := re.ReplaceAllString(string(src), replacement)
	if updated == string(src) {
		fmt.Fprintln(os.Stderr, "version constant not found in ui8kit.go")
		os.Exit(1)
	}

	if err := os.WriteFile("ui8kit.go", []byte(updated), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write ui8kit.go: %v\n", err)
		os.Exit(1)
	}
}
