package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "单元测试",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			fmt.Printf("🔍检查输入参数: %q\n", ctx.Args())
			// if ctx.Args().Len() != 2 {
			// 	fmt.Printf("❌参数错误.  加-help 查看详细用法 \n")
			// 	return nil
			// }
			// var param1 = ctx.Args().Get(0)
			// fmt.Printf("1输入主工程podlock路径: %q\n", param1)
			// var param2 = ctx.Args().Get(1)
			// fmt.Printf("2输入子仓podfile路径: %q\n", param2)

			// /// 拿到主工程依赖
			// fmt.Println("获取依赖...")
			// dependencyMap := podfileLock.FetchDependencies()
			// earth.PrintStrMap(dependencyMap)
			// fmt.Println("\n\n\n")

			// /// 解析子仓podfile
			// fmt.Println("处理子仓podfile...\n Todo: ...")
			// podfile.MakePodfileComefrom(dependencyMap)
			// fmt.Println("\n\n\n")

			fmt.Println("🚀End")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
