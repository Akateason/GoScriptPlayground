package main

import (
	"fmt"
	"goPlay/earth"
	"goPlay/earth/cocoapod/podspec"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "publishPod",
		Usage: "组件发版 \n param(更新版本号第几位 -> 0,1,2)",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("🚀组件发版 \n param(更新版本号第几位 -> 0,1,2) \n")
			fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数
			var param1 = ctx.Args().Get(0)
			if param1 == "" {
				param1 = "2" // 默认index==2, 默认更新最小版本号
				fmt.Printf("不输入参数, 默认输入2更新最小版本号 \n")
			}
			willUpdateVersionIndex := earth.Str2Int(param1)

			newVersion := podspec.UpdateVersion(willUpdateVersionIndex)
			fmt.Printf("%q \n", newVersion)
			gitTag := earth.DeleteQuoteSymbol(newVersion)

			// git提交
			earth.UseCommandLine("git add .")
			earth.UseCommandLine("git commit -m 'publish : " + gitTag + "'")
			earth.UseCommandLine("git tag " + gitTag)
			earth.UseCommandLine("git push gitee master") // 推到gitee 私有库
			earth.UseCommandLine("git push --tags")

			// 发布私有库
			specFileName := podspec.GetSpecFileName()
			cmdLine_updateSpec := "pod repo push XTSpecs " + specFileName + "  --allow-warnings --sources=https://gitee.com/mamba24xtc/xtspecs.git --verbose"
			publishError := earth.UseCommandLine(cmdLine_updateSpec)
			if publishError != nil {
				return publishError
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
