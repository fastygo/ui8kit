#!/usr/bin/env bash
set -euo pipefail

# Usage: ./scripts/release.sh 0.2.0
# Creates a release commit and annotated tag v0.2.0.
#
# Local policy:
# - Always run regular tests.
# - Run race tests only when RELEASE_RUN_RACE=1 and the local toolchain supports CGO.
# - CI remains the authoritative place for go test -race.

VERSION="${1:-}"
RUN_RACE="${RELEASE_RUN_RACE:-0}"

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 0.2.0"
  exit 1
fi

if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$'; then
  echo "Error: version must be in semver format (e.g. 0.2.0), without 'v' prefix"
  exit 1
fi

TAG="v${VERSION}"

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "Error: tag $TAG already exists"
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "Error: working tree is not clean. Commit or stash changes first."
  exit 1
fi

BRANCH="$(git branch --show-current)"
if [ "$BRANCH" != "main" ]; then
  echo "Warning: releasing from branch '$BRANCH' (not main). Continue? [y/N]"
  read -r CONFIRM
  if [ "$CONFIRM" != "y" ] && [ "$CONFIRM" != "Y" ]; then
    echo "Aborted."
    exit 1
  fi
fi

echo "Generating templ files..."
templ generate ./...

echo "Running tests..."
go test ./... -count=1

if [ "$RUN_RACE" = "1" ]; then
  if [ "$(go env CGO_ENABLED)" = "1" ]; then
    echo "Running race tests..."
    go test ./... -race -count=1
  else
    echo "Warning: RELEASE_RUN_RACE=1 but CGO is disabled; skipping local race tests."
    echo "CI will still enforce go test ./... -race on the release tag."
  fi
else
  echo "Skipping local race tests."
  echo "Set RELEASE_RUN_RACE=1 to run them when your environment supports CGO."
fi

echo "Updating Version constant to ${VERSION}..."
go run ./scripts/cmd/set-version "${VERSION}"

git add ui8kit.go
git commit -m "chore: release ${TAG}"

echo "Creating annotated tag ${TAG}..."
git tag -a "$TAG" -m "${TAG}"

echo ""
echo "Done. To publish:"
echo "  git push origin ${BRANCH} --tags"
echo ""
echo "After push:"
echo "  - CI verifies tag/version consistency"
echo "  - CI runs templ generate, go vet, and go test ./... -race"
echo "  - release.yml creates GitHub Release with auto-generated notes"
echo "  - proxy.golang.org indexes ${TAG} within minutes"
echo "  - Users: go get github.com/fastygo/ui8kit@${TAG}"
