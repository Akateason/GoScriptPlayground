/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-08-16 21:07:42
 * @LastEditors: tianchen.xie tianchen.xie@nio.com
 * @LastEditTime: 2023-03-03 18:42:58
 * @FilePath: /boom/boom.go
 * @Description: 单元测试
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"goPlay/earth"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "单元测试",
		Action: func(ctx *cli.Context) error {
			// fmt.Println(ctx.App.Usage)

			// get Arguments 参数
			// fmt.Printf("单元测试~~~args === %q\n", ctx.Args())
			// fmt.Printf("boom! I say %q \n", ctx.Args().Get(0))
			fmt.Println("🚀auto reform view to cell")

			// get cell template first
			template := earth.ReadFileFrom("src/XXXCell.swift.tmp")
			// fmt.Printf(template)

			var files []string
			files, _ = earth.GetAllFilePaths("input", files)
			for i := 0; i < len(files); i++ {
				fileName := files[i]
				if strings.HasSuffix(fileName, ".swift") {
					// get customView's name
					fileName = strings.Replace(fileName, ".swift", "", -1)
					fileName = strings.Replace(fileName, "input/", "", -1)
					fmt.Println(fileName)

					// generate newfile
					content := strings.Replace(template, "XXX", fileName, -1)
					newFileName := "output/" + fileName + "Cell.swift"
					earth.WriteStringToFileFrom(newFileName, content)

				}
			}

			// fmt.Printf(result)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
