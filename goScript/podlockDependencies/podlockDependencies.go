/*
 * @Author: tianchen.xie tianchen.xie@nio.com
 * @Date: 2024-02-22 16:30:00
 * @LastEditors: tianchen.xie tianchen.xie@nio.com
 * @LastEditTime: 2024-03-07 20:36:57
 * @FilePath: /GoScriptPlayground/goScript/podlockDependencies/podlockDependencies.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
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
			fmt.Println("__已输出到dependencies.json")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
