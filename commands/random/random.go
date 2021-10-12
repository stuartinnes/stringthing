package random

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"stringthing/commands"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	RandCMD = &cli.Command{
		Name:    "random",
		Usage:   "generate random strings",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "chars",
				Aliases: []string{"c"},
				Value:   8,
			},
			&cli.IntFlag{
				Name:    "digits",
				Aliases: []string{"d"},
				Value:   8,
			},
			&cli.IntFlag{
				Name:    "symbols",
				Aliases: []string{"s"},
				Value:   8,
			},
		},
		Action: func(c *cli.Context) error {
			var (
				numChars   = c.Int("chars")
				numDigits  = c.Int("digits")
				numSymbols = c.Int("symbols")
			)

			for _, v := range []int{numChars, numDigits, numDigits} {
				if v < 0 {
					return cli.Exit("argument values cannot be negative", commands.EXIT_OTHER)
				}
			}
			return printRandStr(numChars, numDigits, numSymbols)
		},
	}
)

func printRandStr(numChars, numDigits, numSymbols int) error{
	var (
		strLen = numChars + numDigits + numSymbols
		err    error
	)

	if strLen <= 0 {
		return cli.Exit("nothing to do", commands.EXIT_OTHER)
	}

	var (
		randBytes = make([]byte, strLen)
	)

	_, err = rand.Reader.Read(randBytes)
	if err != nil {
		return cli.Exit("unable to generate random sequence", commands.EXIT_OTHER)
	}

	for k, v := range randBytes {
		dict := ""
		switch {
		case k > (numChars+numDigits)-1:
			dict = symbols
		case k > numChars-1:
			dict = digits
		default:
			dict = chars
		}
		randBytes[k] = dict[v%byte(len(dict))]
	}

	mrand.Seed(time.Now().Unix())
	mrand.Shuffle(strLen,
		func(i, j int) {
			randBytes[i], randBytes[j] = randBytes[j], randBytes[i]
		})

	fmt.Println(string(randBytes))
	return nil
}
