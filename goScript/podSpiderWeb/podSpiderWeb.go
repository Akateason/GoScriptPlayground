/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-31 23:06:13
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-31 23:35:37
 * @FilePath: /GoScriptPlayground/goScript/podSpiderWeb/podSpiderWeb.go
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
		Name:  "podSpiderWeb",
		Usage: "查此 sub pod 的 所有间接依赖",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)

			paramSearch := ctx.Args().Get(0)
			fmt.Printf("查此pod的所有间接依赖 , 输入的参数为: %q \n\n\n", paramSearch) // Arguments 参数

			if paramSearch == "" {
				fmt.Printf("💥 没有输入参数, 输入要查的pod \n")
				return nil
			}

			result := podfileLock.FindFather(paramSearch, 0)

			if !result {
				fmt.Printf("👮🏻没有找到" + paramSearch + "的间接依赖~~~~~~\n\n\n")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
