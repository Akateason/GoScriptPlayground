package main

import (
	"fmt"
	"goPlay/earth"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "getSpecVersion",
		Usage: "get podSpec Version",
		Action: func(ctx *cli.Context) error {

			// Arguments 参数
			fmt.Printf("%q \n", ctx.Args().Get(0))
			earth.UseCommandLine("ls")

			var files []string
			files, _ = earth.GetAllFilePaths(".", files)
			fmt.Printf("%q \n", files)

			for i := 0; i < len(files); i++ {
				fileName := files[i]
				// fmt.Print(fileName, "\t")
				if strings.Contains(fileName, ".podspec") {
					fmt.Print(fileName)

					fileContent := earth.ReadFileFrom(fileName)
					fmt.Print(fileContent)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
