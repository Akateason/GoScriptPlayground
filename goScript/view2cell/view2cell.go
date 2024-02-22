/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-06 22:11:29
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-06 22:34:56
 * @FilePath: /GoScriptPlayground/goScript/view2cell/view2cell.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"fmt"
	"goPlay/earth"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "view2cell",
		Usage: "reform view to cell",
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.App.Usage)
			fmt.Printf("输入参数: %q \n", ctx.Args().Get(0)) // Arguments 参数

			template := earth.ReadFileFrom("src/XXXCell.swift.tmp")

			var files []string
			files, _ = earth.GetAllFilePaths("input", files)
			for i := 0; i < len(files); i++ {
				fileName := files[i]
				if strings.HasSuffix(fileName, ".swift") {
					fileName = strings.Replace(fileName, ".swift", "", -1)
					fileName = strings.Replace(fileName, "input/", "", -1)

					content := strings.Replace(template, "XXX", fileName, -1)
					newFileName := "output/" + fileName + "Cell.swift"
					earth.WriteStringToFileFrom(newFileName, content)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
