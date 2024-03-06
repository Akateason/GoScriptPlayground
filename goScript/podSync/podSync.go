package main

import (
	"fmt"
	"goPlay/earth"
	"goPlay/earth/cocoapod/podfile"
	podfileLock "goPlay/earth/cocoapod/podlock"
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

			earth.UseCommandLine("cd " + param1 + ";" + "mkdir -p " + workingFolder + ";" + "cp Podfile.lock " + workingFolder + ";") // ✅get main, podlock
			earth.UseCommandLine("cd " + param2 + ";" + "mkdir -p " + workingFolder + ";" + "cp Podfile " + workingFolder + ";")      // ✅get biz, Podfile

			dependencyMap := podfileLock.FetchDependencies()
			earth.PrintStrMap(dependencyMap)
			fmt.Println()
			if len(dependencyMap) == 0 {
				fmt.Printf("❌ 获取主工程依赖失败, 检查 参数1 \n")
				return nil
			}

			/// 解析子仓podfile
			fmt.Println("🐲处理子仓podfile ing...")
			result := podfile.MakePodfileComefrom(dependencyMap)
			fmt.Println()
			if result {
				fmt.Println("success🚀 \nEnd")
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
