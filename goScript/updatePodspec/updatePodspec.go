package main

import (
	"fmt"
	"goPlay/earth"
	"goPlay/earth/cocoapod/podspec"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "updatePodspec",
		Usage: "update podSpec Version\n param(更新版本号第几位 -> 0,1,2)",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数
			var param1 = ctx.Args().Get(0)
			if param1 == "" {
				param1 = "2" // 默认index==2, 默认更新最小版本号
				fmt.Printf("不输入参数, 默认更新最小版本号 \n")
			}
			willUpdateVersionIndex := earth.Str2Int(param1)

			podspec.UpdateVersion(willUpdateVersionIndex)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
