/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-12 01:07:05
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-14 01:38:33
 * @FilePath: /go/goScript/publishAll/publishAll.go
 * @Description:
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */

package main

import (
	"fmt"
	"go/build"
	"goPlay/earth"
	ggit "goPlay/earth/Git"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "publishAll",
		Usage: "我发布我自己. 编译go为二进制, 安装所有脚本到sender目录. 自动加版本号. param(更新版本号第几位 -> 0,1,2)",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("🚀所有脚本安装与发版 \n param(更新版本号第几位 -> 0,1,2) \n")
			fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数
			var param1 = ctx.Args().Get(0)
			if param1 == "" {
				param1 = "2" // 默认index==2, 默认更新最小版本号
				fmt.Printf("不输入参数, 默认输入2更新最小版本号 \n")
			}

			// 最高tag
			idx := earth.Str2Int(param1)
			tag := ggit.LatestTagVersion()
			tag = earth.UpdateVersionWith(idx, tag)
			fmt.Printf("new version: %q\n\n", tag)
			// git 提交
			earth.UseCommandLine("git add .;git commit -m 'publish " + tag + "';")
			earth.UseCommandLine("git tag " + tag)
			earth.UseCommandLine("git push origin master")
			earth.UseCommandLine("git push gitee master")
			earth.UseCommandLine("git push --tags")

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

			cmdl1 := "cd " + goPath + ";"
			cmdl1 += "cp -r " + ". " + targetPath + ";"
			cmdl1 += "rm -f *;"
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

			// 把publishAll放到根目录. 更新发布脚本.
			cmdl2 := "cp -r sender/publishAll ./"
			earth.UseCommandLine(cmdl2)
			// End
			fmt.Printf("install complete🔥🔥🔥\n\n\n")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
