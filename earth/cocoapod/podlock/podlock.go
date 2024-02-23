/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-09-20 09:12:02
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-29 13:47:25
 * @FilePath: /go/earth/cocoapod/podlock/podlock.go
 * @Description: handle podfile.lock
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package podfileLock

import (
	"fmt"
	"goPlay/earth"
	"strings"
)

// 找.podlock 文件名
func GetFileName() string {
	var files []string
	files, _ = earth.GetAllFilePaths(".", files)
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		if fileName == "./Podfile.lock" {
			return fileName
		}
	}
	return ""
}

// podfile.lock 内容
func FetchContent() string {
	fileName := GetFileName()
	return earth.ReadFileFrom(fileName)
}

// 解析 .lock
func Analysis() []string {
	// fmt.Println("解析开始, 以🐷🐶结尾")
	lockContent := FetchContent()
	sourceList := strings.Split(lockContent, "\n\n")
	// 打印解析结果
	// for _, value := range sourceList {
	// fmt.Println(value + "🐷🐶")
	// }
	return sourceList
}

/**
 * @description: 获取每个pod的数组.
 * @return {sourcelist}
 */
func FetchEverySpecRepos() map[string]interface{} {
	resultMap := make(map[string]interface{})
	var tmpList []string = []string{}
	var currentKey string
	specReposResult := FetchSection("SPEC REPOS:")
	sourceList := strings.Split(specReposResult, "\n")
	for _, v := range sourceList {

		if strings.HasPrefix(v, "SPEC REPOS:") {
			continue
		}

		v = earth.DeleteSpaceSymbol(v)
		v = earth.DeleteQuoteSymbol(v)

		if strings.HasSuffix(v, ":") { // this is repo
			currentKey = v
			tmpList = []string{}
			resultMap[currentKey] = tmpList
		}

		if strings.HasPrefix(v, "-") { // this is pod
			v = strings.TrimPrefix(v, "-")
			tmpList = append(tmpList, v)
			resultMap[currentKey] = tmpList
		}
	}
	return resultMap
}

// 根据name查找对应pod的版本
func CheckPodVersion(podName string) {
	podName = strings.ToLower(podName)
	POD := fetchPOD()
	fmt.Println("🦆🦆🦆搜索🦆🦆🦆" + podName + "的版本号")
	sourceList := strings.Split(POD, "\n")
	for _, v := range sourceList {
		if v[0:3] == "  -" { // 找1级
			var lowerV = strings.ToLower(v)
			if strings.Contains(lowerV, podName) {
				fmt.Println("🦆" + v)
			}
		}
	}
}

// 找间接依赖. return是否找到
func FindFather(podName string, lvl int) bool {
	podName = strings.ToLower(podName)
	POD := fetchPOD()
	sourceList := strings.Split(POD, "\n")
	var tmp1LvlPod string
	var catched = false // podName 是否有间接依赖

	for _, v := range sourceList {
		if v[0:3] == "  -" { // 先存1级
			tmp1LvlPod = cleanPodName(v)
		}

		if v[0:5] == "    -" { // 找2级
			var lowerV = strings.ToLower(cleanPodName(v))

			if lowerV == podName {
				catched = true

				v = cleanPodName(v)
				lvlString := earth.Int2Str(lvl)

				if lvl == 0 {
					fmt.Println("🐱" + v + "\n -> " + lvlString + " " + tmp1LvlPod)
				} else {
					space := space(lvl)
					fmt.Println(space + " -> " + lvlString + " " + tmp1LvlPod)
				}

				FindFather(tmp1LvlPod, lvl+1) // 递归找依赖

				if lvl == 0 {
					fmt.Println("----------") // 只有在第0层打印结束的时候标记结束
				}
			}
		}
	}

	return catched
}

// --------------------------------------------------
// Private
// --------------------------------------------------
// 拿到干净的pod name
func cleanPodName(src string) string {
	src = earth.DeleteSpaceSymbol(src)      // del space
	src = strings.Replace(src, "-", "", -1) // del -
	src = strings.Split(src, "(")[0]        // del ()
	return src
}

// 打印几个space?
func space(number int) string {
	var result = ""
	for i := 0; i < number; i++ {
		result += " "
	}
	return result
}

/**
 * @description: 获取.lock的模块
 * @param {string} name 模块name, 比如PODS:
 * @return {*}
 */
func FetchSection(name string) string {
	analysisResult := Analysis()
	for _, v := range analysisResult {
		if strings.HasPrefix(v, name) {
			return v
		}
	}
	return ""
}

// 获取.lock中 PODS模块
func fetchPOD() string {
	return FetchSection("PODS:")
}

// 获取.lock中 DEPENDENCIES模块
func FetchDependencies() string {
	return FetchSection("DEPENDENCIES:")
}
