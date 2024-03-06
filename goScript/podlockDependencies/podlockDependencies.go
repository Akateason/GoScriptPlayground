package main

import (
	"fmt"
	"goPlay/earth"
	podfileLock "goPlay/earth/cocoapod/podlock"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "podlockDependencies",
		Usage: "解析podlock的依赖, 转成一个json字典",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			fmt.Println("获取依赖...")
			dependencyMap := podfileLock.FetchDependencies()
			earth.PrintStrMap(dependencyMap)
			str, _ := earth.DictToText(dependencyMap)
			earth.WriteStringToFileFrom("dependencies.json", str)
			fmt.Println("已输出到dependencies.json")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
