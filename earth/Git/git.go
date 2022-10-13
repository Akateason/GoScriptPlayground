/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-10-14 00:27:08
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-14 01:18:55
 * @FilePath: /go/earth/git/git.go
 * @Description:
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package ggit

import (
	"goPlay/earth"
)

func LatestTagVersion() string {
	_, tag := earth.ExecuteCommandLine("git describe --tags --abbrev=0")
	return tag
}
