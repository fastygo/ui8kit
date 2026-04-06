#!/usr/bin/env bash
set -euo pipefail

# Usage:
#   ./scripts/preflight.sh
#   PREFLIGHT_RUN_RACE=1 ./scripts/preflight.sh
#
# Checks performed:
# - go generate ./...
# - gofmt -w .
# - go vet ./...
# - go build ./...
# - go test ./... -count=1
# - optional: go test ./... -race -count=1
# - git diff checks to ensure generation/formatting produced no uncommitted changes

RUN_RACE="${PREFLIGHT_RUN_RACE:-0}"

echo "Running go generate..."
go generate ./...

echo "Formatting Go files..."
gofmt -w .

UNFORMATTED="$(gofmt -l .)"
if [ -n "$UNFORMATTED" ]; then
  echo "Error: gofmt still reports unformatted files:"
  echo "$UNFORMATTED"
  exit 1
fi

echo "Running go vet..."
go vet ./...

echo "Building packages..."
go build ./...

echo "Running tests..."
go test ./... -count=1

if [ "$RUN_RACE" = "1" ]; then
  if [ "$(go env CGO_ENABLED)" = "1" ]; then
    echo "Running race tests..."
    go test ./... -race -count=1
  else
    echo "Warning: PREFLIGHT_RUN_RACE=1 but CGO is disabled; skipping local race tests."
    echo "CI will still enforce go test ./... -race."
  fi
else
  echo "Skipping local race tests."
  echo "Set PREFLIGHT_RUN_RACE=1 to run them when your environment supports CGO."
fi

if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "Checking for uncommitted generated or formatted changes..."
  git diff --exit-code
  git diff --cached --exit-code
fi

echo "Preflight OK."
