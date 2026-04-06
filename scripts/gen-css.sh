#!/usr/bin/env bash
set -euo pipefail

# Usage:
#   ./scripts/gen-css.sh
#
# This script generates templ Go files first, then builds
# styles/ui8kit.css safelist classes from UtilityProps mappings and
# literal UtilityProps values found in *_templ.go files.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

cd "${ROOT_DIR}"

echo "Generating templ Go files..."
templ generate

echo "Generating styles/ui8kit.css..."
go run ./scripts/gen-ui8kit-css.go

echo "Done."
