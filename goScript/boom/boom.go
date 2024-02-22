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
	"log"
	"os"

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
			fmt.Println("🚀")

			// fmt.Printf(result)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
