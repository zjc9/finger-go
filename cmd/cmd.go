package cmd

import (
	"context"
	"fmt"
	"github.com/kataras/golog"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"time"
)

var log = golog.Child("[init]")
var Version = "v1.0.0"

func Init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:     "version",
		Aliases:  []string{"v"},
		Usage:    "print the version",
		Category: "[Other Options]",
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:     "help",
		Aliases:  []string{"h"},
		Usage:    "show help",
		Category: "[Other Options]",
	}
	cli.OsExiter = func(code int) {
		fmt.Println(cli.ErrWriter, "", code)
	}
	app := &cli.App{
		Flags:   Flag,
		Action:  Action,
		Version: Version,
		Name:    "指纹识别",
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

func Action(c *cli.Context) error {
	_, cancel := signalCtx()
	defer cancel()

	defer func() {
		var log = golog.Child("[exit]")
		msg := "扫描结束"
		log.Infof(msg)
		time.Sleep(time.Second)
	}()

	fmt.Println(Urls)
	return nil
}

func signalCtx() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		cancel()
	}()
	return ctx, cancel
}
