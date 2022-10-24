/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-12 01:07:05
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-24 23:37:18
 * @FilePath: /go/goScript/publishAll/publishAll.go
 * @Description: 所有脚本发版脚本. 仅供内部使用. [安装到sender]
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */

package main

import (
	"fmt"
	"go/build"
	"goPlay/earth"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "publishAll",
		Usage: "我发布我自己. 编译go为二进制, 安装所有脚本到sender目录. 自动加版本号. \nparam1(根目录绝对路径)\nparam2(更新版本号第几位 -> 0,1,2 可选, 默认=2)",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.UsageText)

			fmt.Printf("检查输入参数: %q \n", ctx.Args()) // Arguments 参数
			fmt.Printf("🚀所有脚本安装与发版 \n param(更新版本号第几位 -> 0,1,2) \n")

			if ctx.Args() == nil {
				fmt.Printf("❌❌❌❌❌ 必须传参. 你不会用 \n")
				return nil
			}
			var param2 = ctx.Args().Get(1)
			if param2 == "" {
				param2 = "2" // 默认index==2, 默认更新最小版本号
				fmt.Printf("不输入参数, 默认输入2更新最小版本号 \n")
			}

			// auto plus tag
			idx := earth.Str2Int(param2)
			_, tag := earth.ExecuteCommandLine("git describe --tags `git rev-list --tags --max-count=1`")
			tag = earth.DeleteNewLine(tag)
			// fmt.Printf("old version was %q\n\n", tag)
			tag = earth.UpdateVersionWith(idx, tag)
			fmt.Printf("new version will be %q\n\n", tag)

			// 开始安装脚本
			fmt.Printf("build All start ...\n\n")
			// get gopath/bin
			goPath := build.Default.GOPATH + "/bin/"

			// target path
			pwd, _ := os.Getwd()

			targetPath := pwd + "/sender/"

			// 对每个goScript文件夹进行 go install
			e1 := earth.UseCommandLine("cd goScript;find . -type d -depth 1 -exec go install {} +")
			if e1 != nil {
				fmt.Printf("❌go scripts 出错\n")
				return e1
			}
			fmt.Printf("go scripts installed\n")

			// 获取所有go脚本Name列表
			e1 = earth.UseCommandLine("cd goScript;find . -type d -depth 1 > ../allgo.txt")
			allgoTxt := earth.ReadFileFrom("allgo.txt")
			allgoList := strings.Split(allgoTxt, "\n")

			cmdl1 := "cd " + goPath + ";"
			for _, v := range allgoList {
				cmdl1 += "cp -r " + v + " " + targetPath + ";"
				cmdl1 += "rm -f " + v + ";"
			}
			e1 = earth.UseCommandLine(cmdl1) // do copy go
			if e1 != nil {
				fmt.Printf("❌go scripts 迁移出错\n")
				return e1
			}

			// 安装shell脚本
			cmdl := "cp -r shell/. " + targetPath
			// fmt.Printf(cmdl + "\n")
			e2 := earth.UseCommandLine(cmdl) // do copy shell
			if e2 != nil {
				fmt.Printf("❌shell scripts 出错\n")
				return e2
			}
			fmt.Printf("shell installed\n")

			cmdl = "cd shell; find *.sh -type f"
			_, allshellTxt := earth.ExecuteCommandLine(cmdl)

			// 把publishAll放到根目录. 更新发布脚本.
			cmdl2 := "cp -r sender/publishAll ./"
			earth.UseCommandLine(cmdl2)
			// End
			fmt.Printf("install complete🔥🔥🔥\n\n\n")

			// readme update
			readme := earth.ReadFileFrom("readme.md")
			readmeList := strings.Split(readme, "# Introduction")
			allgoTxt = strings.Replace(allgoTxt, "./", "", -1)
			readme = readmeList[0] + "# Introduction\n```" + allgoTxt + "\n" + allshellTxt + "```"
			earth.WriteStringToFileFrom("readme.md", readme)

			// git 提交
			earth.UseCommandLine("git add -A .;git commit -m 'publish " + tag + "';")
			earth.UseCommandLine("git tag " + tag)
			earth.UseCommandLine("git push origin master")
			earth.UseCommandLine("git push gitee master")
			earth.UseCommandLine("git push --tags")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
