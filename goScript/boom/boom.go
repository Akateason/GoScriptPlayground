package main

import (
	"fmt"
	"goPlay/earth"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(ctx *cli.Context) error {

			// Arguments 参数
			fmt.Printf("step1 boom! I say %q \n", ctx.Args().Get(0))

			earth.UseCommandLine("cd ../../..;ls -l")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
