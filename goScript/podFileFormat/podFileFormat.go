/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-09-19 23:00:20
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-29 13:48:02
 * @FilePath: /go/goScript/podFileFormat/podFileFormat.go
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
		Usage: "podFileFormat...",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("格式化Podfile \n")
			//fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			newContent := podfile.ExportNewPodfile()
			newPath := "format_副本_pod_file"
			earth.UseCommandLine("touch " + newPath)
			earth.WriteStringToFileFrom(newPath, newContent)

			fmt.Printf("\n\n\n🐂🐴\n\n\n格式化成功, 查看format_副本_pod_file \n")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
