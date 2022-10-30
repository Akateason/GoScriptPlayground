/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-29 10:49:57
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-30 22:01:48
 * @FilePath: /go/goScript/pod2Local/pod2Local.go
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
		Name:  "pod2Local",
		Usage: "pod2Local \nPodfile切本地. 请在Podfile同级目录下配置CONFIG_pod2Local为json文件. \n",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			// fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			var configPath = "CONFIG_pod2Local"
			earth.IfNoFileToCreate(configPath)
			tmpStr := earth.ReadFileFrom(configPath)
			if len(tmpStr) == 0 {
				fmt.Println("❌ 没有配置json到 CONFIG_pod2Local.")
				fmt.Println("请配置.")
				earth.WriteStringToFileFrom(configPath, "{\"pod_name\":\":path=>'../your_path'\"}")

				return nil
			}

			tmpMap := earth.JsonStrToMap(tmpStr)
			fmt.Printf("%q\n", tmpMap)

			_ = podfile.ConfigPodfileWithMap(tmpMap)
			// fmt.Printf("%q", whatMap)
			// fmt.Printf("\n\n\n🐂🐴\n\n\n格式化成功, 查看format_副本_pod_file \n")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
