/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-11 01:03:42
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-15 20:10:20
 * @FilePath: /go/goScript/podGitAddress/podGitAddress.go
 * @Description: 查pod远程仓库地址
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
		Name:  "podGitAddress",
		Usage: "查所有pod的远程仓库地址",
		Action: func(ctx *cli.Context) error {

			fmt.Println("查所有pod的远程仓库地址")
			fmt.Println("start ... ")

			resultlist := podfileLock.FetchEverySpecRepos()

			for _, v := range resultlist {
				cmlStr := "pod search " + v + " > tmp.txt;"
				cmlStr += "awk '/->/ {print $0; exit; }' tmp.txt;"
				cmlStr += "awk '/Source/ {print $3; exit; }' tmp.txt"
				earth.UseCommandLine(cmlStr)
				fmt.Println("========================================")
			}

			earth.UseCommandLine("rm -f tmp.txt")
			fmt.Println("[Done]")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
