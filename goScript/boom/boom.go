/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-08-16 21:07:42
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-11 00:58:50
 * @FilePath: /go/goScript/boom/boom.go
 * @Description: 单元测试
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"goPlay/earth"
	podfileLock "goPlay/earth/cocoapod/podlock"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance,  unit test, playground",
		Action: func(ctx *cli.Context) error {
			fmt.Println("boom! I say~")
			// get Arguments 参数
			fmt.Printf("单元测试~~~args === %q\n", ctx.Args())
			// fmt.Printf("boom! I say %q \n", ctx.Args().Get(0))
			// earth.UseCommandLine("cd ../../..;ls -l")

			fmt.Println("start ... ")
			// earth.UseCommandLine("pod repo update")
			resultlist := podfileLock.FetchEverySpecRepos()

			for _, v := range resultlist {
				fmt.Printf(v + "👉🏻")
				cmlStr := "pod search " + v + " > tmp.txt;"
				cmlStr += "awk '/Source/ {print $3; exit; }' tmp.txt"
				earth.UseCommandLine(cmlStr)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
