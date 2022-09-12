package main

import (
	"goPlay/earth/cocoapod/podspec"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "getSpecVersion",
		Usage: "get podSpec Version",
		Action: func(ctx *cli.Context) error {
			// Arguments 参数
			//fmt.Printf("%q \n", ctx.Args().Get(0))

			// var _ = podspec.GetPodSpecContent()
			var _ = podspec.GetVersion()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
