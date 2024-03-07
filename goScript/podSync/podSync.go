/*
 * @Author: tianchen.xie tianchen.xie@nio.com
 * @Date: 2024-02-22 16:30:00
 * @LastEditors: tianchen.xie tianchen.xie@nio.com
 * @LastEditTime: 2024-03-07 14:20:16
 * @FilePath: /GoScriptPlayground/goScript/podSync/podSync.go
 * @Description: podSync
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
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
		Name:  "podSync",
		Usage: "同步主工程的PodFile到子仓. \n务必输入参数 \n1输入主工程podlock路径 \n2输入子仓podfile路径 \n用空格衔接",
		Action: func(ctx *cli.Context) error {
			fmt.Println("🐲start .")
			fmt.Println(ctx.App.Usage) // desc

			// CHECK params
			fmt.Printf("🔍检查输入参数: %q\n", ctx.Args())
			if ctx.Args().Len() != 2 {
				fmt.Printf("❌参数错误.  加-help 查看详细用法 \n")
				return nil
			}
			var param1 = ctx.Args().Get(0)
			fmt.Printf("1输入主工程podlock路径: %q\n", param1)
			if len(param1) == 0 {
				fmt.Printf("❌参数错误.  加-help 查看详细用法 \n")
				return nil
			}
			var param2 = ctx.Args().Get(1)
			if len(param2) == 0 {
				fmt.Printf("❌参数错误.  加-help 查看详细用法 \n")
				return nil
			}
			fmt.Printf("2输入子仓podfile路径: %q\n", param2)

			// new 工作区
			workingFolder := "~/Desktop/workingspace"
			/// 拿到主工程依赖
			fmt.Println("🐲获取主工程依赖 ing...")

			earth.UseCommandLine("cd " + param1 + ";" + "mkdir -p " + workingFolder + ";" + "cp Podfile.lock " + workingFolder + ";")          // ✅get main, podlock
			earth.UseCommandLine("cd " + param2 + ";" + "mkdir -p " + workingFolder + ";" + "cp Podfile " + workingFolder + ";podFileFormat;") // ✅get biz, Podfile, and format podfile

			earth.UseCommandLine("cd " + workingFolder + ";" + "podlockDependencies" + ";") // ✅find dependency

			_, jsonDependency := earth.ExecuteCommandLine("cd " + workingFolder + ";str=$(cat dependencies.json);echo $str;") // ✅fetch dependency.json
			dependencyMap, _ := earth.TextToDict(jsonDependency)
			// earth.PrintStrMap(dependencyMap)
			fmt.Println()
			if len(dependencyMap) == 0 {
				fmt.Printf("❌ 获取主工程依赖失败, 检查 参数1 \n")
				return nil
			}

			/// 解析子仓podfile
			fmt.Println("🐲处理子仓podfile ing...")

			_, absoluteWorkingspacePath := earth.TransLinuxPathToAbsolutePath(workingFolder) // ✅golang不能识别波浪号路径"~/xxx", 转成绝对路径

			podfileContent := earth.ReadFileFrom(absoluteWorkingspacePath + "/" + "Podfile") // ✅ get podfile content from workspace

			isSuccess, result := podfile.MakePodfileComefrom(dependencyMap, podfileContent) // ✅拿到新Podfile整合结果

			earth.WriteStringToFileFrom(param2+"/Podfile", result)
			fmt.Println()
			if isSuccess {
				fmt.Println("success🚀🚀🚀 \nEnd")
			} else {
				fmt.Printf("❌ 解析子仓podfile失败, 检查 参数2 \n")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
