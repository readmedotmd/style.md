#!/bin/bash
set -euo pipefail

# DO Functions: set up Go module with local replace for the shared generate package.
# The CI workflow copies generate/ into this directory before deploy.

if [ ! -f generate/go.mod ]; then
  cd generate
  go mod init github.com/readmedotmd/style.md/generate
  cd ..
fi

if [ ! -f go.mod ]; then
  go mod init exec
  go mod edit -replace github.com/readmedotmd/style.md/generate=./generate
  go mod tidy
fi
