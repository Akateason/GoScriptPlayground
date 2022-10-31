/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-31 23:06:13
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-31 23:35:21
 * @FilePath: /GoScriptPlayground/goScript/podV/podV.go
 * @Description:
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
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
		Name:  "podV",
		Usage: "找pod版本号, 输入搜索版本",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			fmt.Printf("输入的参数为: %q \n", ctx.Args()) // Arguments 参数
			paramSearch := ctx.Args().Get(0)
			if paramSearch == "" {
				fmt.Printf("❎ 你没有输入参数, 输入要查的pod \n")
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
