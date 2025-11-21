/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-09-19 23:00:20
 * @LastEditors: tianchen.xie tianchen.xie@nio.com
 * @LastEditTime: 2024-03-07 20:21:19
 * @FilePath: /GoScriptPlayground/goScript/podFileFormat/podFileFormat.go
 * @Description:
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"goPlay/earth"
	"goPlay/earth/cocoapod/podfile"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "podFileFormat",
		Usage: "格式化Podfile",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			fmt.Printf("格式化Podfile \n")
			//fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数
			tmpstr := podfile.FetchContent()
			cleared := podfile.RemoveAllAnnoation(tmpstr)
			fmt.Println("清理注释...\n" + cleared)
			newContent := podfile.ExportFomatedPodfile(cleared)
			earth.WriteStringToFileFrom("Podfile", newContent)
			fmt.Printf("\n\n\n\n格式化成功, 查看Podfile \n")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
