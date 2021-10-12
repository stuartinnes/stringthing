package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"stringthing/commands/is"
	"stringthing/commands/random"
	"stringthing/commands/uuid"
)

func main() {
	app := cli.App{
		Usage: "A handy CLI string toolkit ",
		Commands: []*cli.Command{
			is.FormatCMD,
			random.RandCMD,
			uuid.UUIDCMD,
		},
	}
	app.Run(os.Args)
}
