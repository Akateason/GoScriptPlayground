package main

import (
	"fmt"
	podfileLock "goPlay/earth/cocoapod/podlock"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "podV",
		Usage: "找pod版本号",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("查pod版本, 输入的参数为: %q \n", ctx.Args()) // Arguments 参数
			paramSearch := ctx.Args().Get(0)
			if paramSearch == "" {
				fmt.Printf("❎ 你没有输入参数, 输入要查的pod \n")
				return nil
			}

			podfileLock.CheckPodVersion(paramSearch)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
