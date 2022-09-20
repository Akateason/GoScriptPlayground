package main

import (
	podfileLock "goPlay/earth/cocoapod/podlock"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "podLockAnalysis",
		Usage: "podfile.lock Analysis...",
		Action: func(ctx *cli.Context) error {
			// fmt.Printf("格式化Podfile \n")
			//fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			podfileLock.Analysis()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
