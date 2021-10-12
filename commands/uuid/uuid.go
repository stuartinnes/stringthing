package uuid

import (
	"fmt"
	"github.com/google/uuid"
	shorty "github.com/lithammer/shortuuid"
	"github.com/urfave/cli/v2"
)

var (
	UUIDCMD = &cli.Command{
		Name:    "uuid",
		Usage:   "generate V4 uuids",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "short",
				Aliases: []string{"s"},
				Usage:   "short form UUID",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("short") {
				fmt.Println(shorty.New())
			} else {
				fmt.Println(uuid.New().String())
			}
			return nil
		},
	}
)
