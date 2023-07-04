package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
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

func UrlHandlerAction(ctx *cli.Context, v string) error {
	for i := 0; i < ctx.NArg(); i++ {
		fmt.Printf("%d: %s\n", i+1, ctx.Args().Get(i))
	}
	//fmt.Println("ctx.narg,", ctx.NArg())
	if strings.Contains(v, ",") {
		urls := strings.Split(v, ",")
		Urls = make([]string, len(urls))
		copy(Urls, urls)
		for i, url := range Urls {
			fmt.Println(i, url)
		}
	} else if ctx.NArg() > 0 {

	}

	return nil

}
func OutPutAction(ctx *cli.Context, v string) error {
	fmt.Println("output", v)
	return nil

}
