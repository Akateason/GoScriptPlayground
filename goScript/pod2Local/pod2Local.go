/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-09 01:40:43
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-10 23:24:02
 * @FilePath: /go/goScript/pod2Local/pod2Local.go
 * @Description: WIP 切换本地
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pod2Local",
		Usage: "Podfile切本地.",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("Podfile切本地. \n")
			fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			// TODO:

			// newContent := podfile.ExportNewPodfile()
			// newPath := "format_副本_pod_file"
			// earth.UseCommandLine("touch " + newPath)
			// earth.WriteStringToFileFrom(newPath, newContent)
			// fmt.Printf("\n\n\n🐂🐴\n\n\n格式化成功, 查看format_副本_pod_file \n")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
