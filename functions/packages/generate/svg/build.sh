#!/bin/bash
set -euo pipefail

# DO Functions builds each function in isolation. The workflow copies the
# shared generate/ package into this directory before deploy. Set up the
# module with a local replace directive so it resolves.

# Create a go.mod for the copied generate package
cd generate
go mod init github.com/readmedotmd/style.md/generate
go mod edit -go=1.23
cd ..

go mod init exec
go mod edit -go=1.23
go mod edit -replace github.com/readmedotmd/style.md/generate=./generate
go mod tidy
