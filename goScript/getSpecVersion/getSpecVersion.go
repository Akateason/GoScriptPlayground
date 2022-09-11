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
		Name:  "getSpecVersion",
		Usage: "get podSpec Version",
		Action: func(ctx *cli.Context) error {
			// Arguments 参数
			fmt.Printf("%q \n", ctx.Args().Get(0))
			// earth.UseCommandLine("ls")

			var _ = podspec.GetPodSpecContent()
			var 


			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
