package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var Flag = []cli.Flag{
	&cli.StringFlag{
		Name:     "o",
		Usage:    "output",
		Action:   OutPutAction,
		Category: "[\x00Output Options]",
	},
	&cli.StringFlag{
		Name:     "u",
		Usage:    "url",
		Action:   UrlHandlerAction,
		Category: "[\x00Output Options]",
	},
	&cli.StringFlag{
		Name:     "f",
		Usage:    "file",
		Action:   OutPutAction,
		Category: "[\x00Output Options]",
	},
}

// cli.VersionFlag = &cli.BoolFlag{
// Name:     "version",
// Aliases:  []string{"v"},
// Usage:    "print the version",
// Category: "[Other Options]",
// }
// cli.HelpFlag = &cli.BoolFlag{
// Name:     "help",
// Aliases:  []string{"h"},
// Usage:    "show help",
// Category: "[Other Options]",
// }
func UrlHandlerAction(ctx *cli.Context, v string) error {
	fmt.Println(ctx.Args().First(), v)

	return nil

}
func OutPutAction(ctx *cli.Context, v string) error {
	fmt.Println("123", v)
	return nil

}
