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
		Name:  "checkPod",
		Usage: "podfile.lock Analysis之后. 搜索pod版本号..",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("查pod版本, 输入参数: %q \n", ctx.Args()) // Arguments 参数
			paramSearch := ctx.Args().Get(0)
			if paramSearch == "" {
				fmt.Printf("你没输入参数, 我挂了 \n")
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
