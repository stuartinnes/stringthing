package is

import (
	"github.com/asaskevich/govalidator"
	"github.com/urfave/cli/v2"
	"strconv"
	"stringthing/commands"
)

var (
	FormatCMD = &cli.Command{
		Name:        "is",
		Usage:       "string format checks",
		Subcommands: subCommands,
	}
)

func doFormatAction(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.Exit("nothing to test", commands.EXIT_NOK)
	}
	return formatResult(govalidator.TagMap[c.Command.Name](c.Args().First()), c.Bool("verbose"))
}

func formatResult(result, verbose bool) error {
	resultStr := strconv.FormatBool(result)
	if verbose {
		if !result {
			return cli.Exit(resultStr, commands.EXIT_NOK)
		}
		return cli.Exit(resultStr, 0)
	}

	if !result {
		return cli.Exit("", commands.EXIT_NOK)
	}
	return cli.Exit("", 0)
}
