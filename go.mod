module github.com/SirMetathyst/atom

go 1.13

require (
	github.com/SirMetathyst/atomkit v0.2.0
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli/v2 v2.0.0
)

replace github.com/SirMetathyst/atomkit v0.2.0 => ../kit/atomkit
