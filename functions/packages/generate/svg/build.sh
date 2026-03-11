#!/bin/bash
set -euo pipefail

# DO Functions builds each function in isolation. The workflow copies the
# shared generate/ package into this directory before deploy. Set up the
# module with a local replace directive so it resolves.

go mod init exec
go mod edit -replace github.com/readmedotmd/style.md/generate=./generate
go mod tidy
