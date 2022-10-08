package main

import (
	"fmt"
	"go/build"
	"goPlay/earth"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "installAll",
		Usage: "⚠️⚠️⚠️ 更新并安装所有脚本集. ⚠️⚠️⚠️ 在[根目录]执行./installAll即刻体验",
		Action: func(ctx *cli.Context) error {

			fmt.Printf("🕛install All start ...\n\n")

			// 拉新
			earth.UseCommandLine("git pull")

			// 对每个goScript文件夹进行 go install
			e1 := earth.UseCommandLine("cd goScript;find . -type d -depth 1 -exec go install {} +")
			if e1 != nil {
				fmt.Printf("❌go scripts 出错\n")
				return e1
			}
			fmt.Printf("go scripts installed\n")

			// 安装shell脚本
			toPath := build.Default.GOPATH
			toPath += "/bin/" // get gopath/bin
			//fmt.Printf("\ngoPath: %q \n", toPath)
			cmdl := "cp -r shell/. " + toPath
			//fmt.Printf(cmdl + "\n")
			e2 := earth.UseCommandLine(cmdl) // do copy shell
			if e2 != nil {
				fmt.Printf("❌shell scripts 出错\n")
				return e2
			}
			fmt.Printf("shell installed\n")

			// End
			fmt.Printf("install complete🔥🔥🔥\n\n\n")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
