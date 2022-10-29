/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-09-19 23:07:46
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2022-10-29 15:00:45
 * @FilePath: /go/earth/cocoapod/podfile/podfile.go
 * @Description: podfile工具
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package podfile

import (
	"fmt"
	"goPlay/earth"
	"strings"
)

// 获取对应PodFile的文件名
func GetPodfileFileName() string {
	var files []string
	files, _ = earth.GetAllFilePaths(".", files)
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		if strings.Contains(fileName, "Podfile") {
			return fileName
		}
	}
	return ""
}

// 获取Podfile内容
func FetchContent() string {
	fileName := GetPodfileFileName()
	return earth.ReadFileFrom(fileName)
}

// 解析Podfile. 分组
// 1.
// 忽略 纯\n
// 忽略 target do ... end  嵌套 忽略各种关键字.
// 忽略 #注释
// 2.
// 根据pod内容分组
func Analysis(needPrint bool) []string {
	fmt.Println("解析开始, 以🐷🐶结尾")
	var resultList []string
	podfileContent := FetchContent()
	sourceList := strings.Split(podfileContent, "\n")
	for _, value := range sourceList {
		if isAllWhiteSpace(value) {
			continue
		}
		if isAnnoation(value) {
			continue
		}
		if isTargetDoEnd(value) {
			continue
		}
		// 打印原始解析
		// fmt.Println("Index =", index, "Value =", value)
		if firstWordIsPod(value) {
			resultList = append(resultList, value)
		} else {
			lastValue := resultList[len(resultList)-1]
			lastValue += value
			resultList[len(resultList)-1] = lastValue
		}
	}

	// 打印解析结果
	if needPrint {
		for _, value := range resultList {
			fmt.Println(value + "🐷🐶")
		}
	}

	return resultList
}

// 3.
// podFileFormat 导出新Podfile
func ExportNewPodfile() string {
	fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	oldPodfile := FetchContent()
	resultList := Analysis(true)
	for _, value := range resultList {
		oldStr := findSourceLineWith(value, oldPodfile)
		fmt.Println("搜索" + value + "\n")
		fmt.Println("得出" + oldStr + "\n--------\n")
		if len(oldStr) > 0 {
			oldPodfile = strings.Replace(oldPodfile, oldStr, value, 1)
		}
	}
	fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	return oldPodfile
}

// podKey
// pod name
const kPodName = "podName"

// 上个版本. 用来reset
const kOriginContent = "originContent"

// 本地路径
const kLocalPath = "localPath"

// // 远程路径
// // 1. git仓库信息
// const kGitRemotePath_andTag = "remotePath+tag"
// const kGitRemotePath_andBranch = "remotePath+branch"
// const kGitRemotePath_andCommit = "remotePath+commit"

// // 2. pod版本号
// const kVerison = "version"

// 嵌套字典 声明
// type Type_str_str_map map[string]string

// pod来源状态.
// 只能是 kLocalPath,
const kPodResourceState = "state"

/*
*

  - @description: 将pod按照本地配置进行处理. 并返回

  - @param localPathMap 一个字典套字典, 映射表. 可以是任何pod后的内容.
    localPathMap =
    [podName : [originContent:string!, localPath:string?, remotePath:string?, branch:string?, commitHash:string?]]

  - @return { HistryMapMap返回保留更改之前的信息. }
*/
func ConfigPodfileWithMap(soureMap map[string]string) map[string]string {
	fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	oldPodfile := FetchContent()
	analsisList := Analysis(false)
	// loop source map
	for podNameKey, contentValue := range soureMap {
		// fmt.Println(podNameKey)
		// fmt.Println(contentValue)

		for _, podValue := range analsisList {
			if strings.Contains(podValue, "\""+podNameKey+"\"") ||
				strings.Contains(podValue, "'"+podNameKey+"'") {
				// podfile is matched !
				fmt.Println(podNameKey + " - is matched !🐶")
				fmt.Println("--- " + podValue)

				oldStr := findSourceLineWith(podValue, oldPodfile)
				fmt.Println("搜索" + podValue)
				fmt.Println("得出" + oldStr + "\n--------\n")
				if len(oldStr) > 0 {
					newPodValue := makeNewPodItemToLocalPath(podValue, contentValue) // to local path
					oldPodfile = strings.Replace(oldPodfile, oldStr, newPodValue, 1)
				}

				// contentMap[kOriginContent] = podValue
			}
		}
	}

	fmt.Println(oldPodfile)
	fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	return soureMap
}

// -------------------------------------------------- //
// -- Private

/*
*
  - @description: 制作 拼接本地podfile的单行.
  - @param {podItemSource} 类似
    pod "MPDebugTools",
    :subspecs => ["Vehicle", "CNLink", "CNAccount", "Review","AntiFraud"],
    :configurations => ['Debug','Test'],
    :git=>"git@git.nevint.com:ios_dd/mpdebugtools.git", :commit=>'2fada45c9d31d8fcb2669773d3dcd747d74deb8c'
  - @param {*} appendValue 逗号后面的东西. "../../snapkit"
  - @return
    pod "MPDebugTools",
    :subspecs => ["Vehicle", "CNLink", "CNAccount", "Review","AntiFraud"],
    :configurations => ['Debug','Test'], :path=>"../../snapkit"
*/
func makeNewPodItemToLocalPath(podItemSource string, appendValue string) string {
	var podPrefix string
	if strings.Contains(podItemSource, ",") {
		podItems := strings.Split(podItemSource, ",")
		var newItems []string
		for _, maohaoItem := range podItems { //
			if isAbsolutelyNeedItem(maohaoItem) {
				newItems = append(newItems, maohaoItem)
			}
		}
		podPrefix = strings.Join(newItems, ",")
	} else {
		podPrefix = podItemSource
	}
	return podPrefix + ", :path=>\"" + appendValue + "\"\n"
}

// 切pod元素.  判断是否应该保留逗号分割的元素
func isAbsolutelyNeedItem(source string) bool {
	if strings.Contains(source, "pod ") {
		return true
	}
	source = earth.DeleteNewLine(source)
	source = earth.DeleteSpaceSymbol(source)
	if strings.Contains(source, ":subspecs") {
		return true
	}
	if strings.Contains(source, ":configurations") {
		return true
	}
	if strings.Contains(source, ":") { // 保留其他带冒号item
		return false
	}
	if (strings.HasPrefix(source, "'") && strings.HasSuffix(source, "'")) ||
		(strings.HasPrefix(source, "\"") && strings.HasSuffix(source, "\"")) {
		// 版本号去掉
		return false
	}
	return true
}

// 字符串全部都是空格?
func isAllWhiteSpace(source string) bool {
	source = earth.DeleteSpaceSymbol(source)
	return source == ""
}

// 字符串是注释?
func isAnnoation(source string) bool {
	source = earth.DeleteSpaceSymbol(source)
	if len(source) > 0 {
		if source[0:1] == "#" {
			return true
		}
	}
	return false
}

// 字符串是 "target do, end, use_frameworks" 等Podfile中无关的关键字?
func isTargetDoEnd(source string) bool {
	if strings.Contains(source, "target") &&
		strings.Contains(source, "do") {
		return true
	}
	source = earth.DeleteSpaceSymbol(source)
	if source == "end" {
		return true
	}
	if strings.Contains(source, "use_frameworks") {
		return true
	}
	if strings.Contains(source, "source") {
		return true
	}
	if strings.Contains(source, "platform") {
		return true
	}
	if strings.Contains(source, "post_install") {
		return true
	}
	if strings.Contains(source, "config.") {
		return true
	}
	if strings.Contains(source, "installer.") {
		return true
	}

	return false
}

// 第一个词是pod?
func firstWordIsPod(source string) bool {
	source = earth.DeleteSpaceSymbol(source)
	if len(source) >= 3 {
		if source[0:3] == "pod" {
			return true
		}
	}
	return false
}

// 找出value对应在source中的原文string
func findSourceLineWith(value string, podfileSource string) string {
	//fmt.Println("搜索" + value + "\n")
	var resultString string = ""
	if strings.Contains(value, ",") {
		// 有条件的pod, 例如像pod 'XTFMDB', :path=>'../XTFMDB'
		lineSourceList := strings.Split(podfileSource, "\n")
		firstPodPark := strings.Split(value, ",")[0]

		var theIndex = -2
		for index, v := range lineSourceList {
			if strings.Contains(v, firstPodPark) {
				theIndex = index
				resultString = v
			} else if theIndex+1 == index {
				if firstWordIsPod(v) || isTargetDoEnd(v) {
					return resultString
				} else {
					theIndex++
					resultString += "\n"
					resultString += v
				}
			}
		}
	} else {
		// 纯 pod "file". 直接返回
		resultString = value
	}
	return resultString
}
