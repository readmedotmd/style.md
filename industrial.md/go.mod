module github.com/readmedotmd/industrial.md

go 1.23.6

require (
	github.com/readmedotmd/core.md v0.0.0
	github.com/readmedotmd/gui.md v0.0.0-20260310051437-76abebc27345
)

replace (
	github.com/readmedotmd/core.md => ../core.md
	github.com/readmedotmd/gui.md => ../../gui.md
)
