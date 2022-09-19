package main

import (
	"goPlay/earth"
	"goPlay/earth/cocoapod/podfile"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "podFileFormat",
		Usage: "podFileFormat...",
		Action: func(ctx *cli.Context) error {
			//fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			// podfile.Analysis()
			newContent := podfile.ExportNewPodfile()
			newPath := "format_副本_pod_file"
			earth.UseCommandLine("touch " + newPath)
			earth.WriteStringToFileFrom(newPath, newContent)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
