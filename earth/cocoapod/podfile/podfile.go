/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-09-19 23:07:46
 * @LastEditors: tianchen.xie tianchen.xie@nio.com
 * @LastEditTime: 2024-03-07 15:47:15
 * @FilePath: /podSync/Users/tianchen.xie/Documents/GoScriptPlayground/earth/cocoapod/podfile/podfile.go
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
		if fileName == "./Podfile" {
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
func AnalysisLocal(needPrint bool) []string {
	podfileContent := FetchContent()
	return Analysis(needPrint, podfileContent)
}

func Analysis(needPrint bool, podfileContent string) []string {
	var resultList []string
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
			// fmt.Println("%q", len(resultList))
			// fmt.Println(resultList)
			if len(resultList) > 0 {
				lastValue := resultList[len(resultList)-1]
				lastValue += value
				resultList[len(resultList)-1] = lastValue
			}
		}
	}

	// 打印解析结果
	if needPrint {
		fmt.Println("解析开始")
		for _, value := range resultList {
			fmt.Println(value + "🐷🐶")
		}
	}

	return resultList
}

// 3.
// podFileFormat 导出新Podfile
func ExportFomatedPodfile() string {
	fmt.Println(" podfileformat🐲🐲🐲🐲🐲🐲🐲 ")
	oldPodfile := FetchContent()
	resultList := AnalysisLocal(true)
	for _, value := range resultList {
		oldStr := findSourceLineWith(value, oldPodfile)
		// fmt.Println("搜索" + value + "\n")
		// fmt.Println("得出" + oldStr + "\n--------\n")
		if len(oldStr) > 0 {
			clearedValue := earth.DeleteSpaceSymbol(value)
			oldPodfile = strings.Replace(oldPodfile, oldStr, clearedValue, 1)
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
  - @return {
    HistryMapMap返回保留更改之前的信息.
    }
*/
// @deprecated: This method will be removed in future releases
func Pod2LocalConfigPodfileWithMap(soureMap map[string]interface{}) map[string]interface{} {
	fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	newPodfile := FetchContent()

	analsisList := AnalysisLocal(false)

	var historyMap map[string]interface{} = make(map[string]interface{})

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

				oldStr := findSourceLineWith(podValue, newPodfile)

				fmt.Println("搜索" + podValue)
				fmt.Println("得出" + oldStr + "\n--------\n")

				if len(oldStr) > 0 {
					newPodValue := makeNewPodItemToLocalPath(podValue, contentValue.(string)) // to local path
					newPodfile = strings.Replace(newPodfile, oldStr, newPodValue, 1)

					historyMap[podNameKey] = oldStr // 记录上一次的历史
				}

			}
		}
	}

	if len(historyMap) > 0 {
		fmt.Println(newPodfile) // 新podfile
		fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")

		// 删除这段逻辑 . 没必要记录. 有git
		// output newpodfile, and save old podfile
		// oldPodfile := FetchContent()
		// earth.UseCommandLine("touch " + "oldPodfile")
		// earth.WriteStringToFileFrom("oldPodfile", oldPodfile)

		earth.WriteStringToFileFrom("Podfile", newPodfile)

		// for k, v := range historyMap {
		// 	fmt.Printf("key: %q\n", k)
		// 	fmt.Printf("val: %q\n", v)
		// }

		// 删除这段逻辑 . 没必要记录. 有git
		// make history
		// timeStr := time.Now().Format("20220101_11:11:01")
		// newHistroyPath := "before_pod2Local" + timeStr
		// earth.UseCommandLine("touch " + newHistroyPath)
		// jsonStr := earth.MapToJsonStr(historyMap)
		// earth.WriteStringToFileFrom(newHistroyPath, jsonStr)

	} else {
		fmt.Println("pod2local isMatched. or failed. ❌")
	}

	return historyMap
}

/**
 * @description: 通用做Podfile方法, 统一改来源
 * @param {map[string]string} soureMap 一个字典套字典, 映射表. 可以是任何pod后的内容.
 * @param {podfileContent} 来源podfile 文本内容
 * @return {success, result新Podfile文本}
 */
func MakePodfileComefrom(sourceMap map[string]string, podfileContent string) (bool, string) {
	// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 1")
	// fmt.Println(soureMap)
	// podfileContent := FetchContent(), // 不用本地路径下了, 做成参数进来.

	analsisList := Analysis(false, podfileContent)

	for _, podValue := range analsisList {
		// podValue = earth.DeleteSpaceSymbol(podValue)
		// podValue = earth.DeleteNewLine(podValue)
		podName := getOneLinePodName(podValue)

		contentValue, ok := sourceMap[podName]
		// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲1.1=" + podName)

		if ok {
			// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 2")
			// fmt.Println(podName + " - is matched !🐶")
			// fmt.Println("---> " + podValue)

			originStrFromOldContent := findSourceLineWith(podValue, podfileContent)
			if strings.Contains(originStrFromOldContent, ":path") { // 如果指向本地, 则忽略覆盖
				fmt.Println(podName + "子仓指向本地, 忽略")
				continue
			}

			// fmt.Println("🐲🐲🐲搜索2.11🐲" + podValue)
			// fmt.Println("🐲🐲🐲搜索2.12🐲" + originStrFromOldContent)
			if len(originStrFromOldContent) > 0 {
				var podPrefix string
				if strings.Contains(podValue, ",") {
					// fmt.Println("🐲🐲🐲2.13🐲" + originStrFromOldContent)
					clearedPodValue := earth.DeleteSpaceSymbol(podValue)
					podItems := strings.Split(clearedPodValue, ",:") //拆分组
					// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 2.2")
					// fmt.Println(podItems)

					var newItems []string
					for _, maohaoItem := range podItems { //
						// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 2.3 冒号" + maohaoItem)
						if strings.HasPrefix(maohaoItem, "pod") &&
							strings.Contains(maohaoItem, ",") {
							maohaoItem = strings.Split(maohaoItem, ",")[0]
						}
						if isAbsolutelyNeedItem(maohaoItem) {
							newItems = append(newItems, maohaoItem)
						}
					}
					podPrefix = strings.Join(newItems, ",")
				} else {
					podPrefix = podValue
				}
				podPrefix = earth.DeleteSpaceSymbol(podPrefix) // del space
				podPrefix = earth.DeleteNewLine(podPrefix)     // del \n
				// fmt.Println("得" + podPrefix)

				if !strings.HasPrefix(contentValue, ",") {
					contentValue = "," + contentValue
				}
				newPodValue := podPrefix + contentValue
				newPodValue = earth.DeleteSpaceSymbol(newPodValue)
				podfileContent = strings.Replace(podfileContent, originStrFromOldContent, newPodValue, 1)

				// fmt.Println("出" + contentValue)
				// fmt.Println("得出" + newPodValue + "\n--------\n")
			}
		}
	}

	fmt.Println(podfileContent) // 新podfile
	if len(podfileContent) > 0 {
		// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲3 ")
		return true, podfileContent
	}
	return false, ""
}

// -------------------------------------------------- //
// -------------------------------------------------- //
// -------------------------------------------------- //
// -- Private
// -------------------------------------------------- //
// -------------------------------------------------- //
// -------------------------------------------------- //

/**
 * @description: 拿到这行的pod名字
 * @param {string} oneLine
 * @return {*}
 */
func getOneLinePodName(oneLine string) string {
	if strings.HasPrefix(strings.TrimSpace(oneLine), "pod") {
		oneLine = earth.DeleteSpaceSymbol(oneLine)
		parts := strings.TrimPrefix(oneLine, "pod")
		parts = strings.Split(parts, ",")[0]
		if len(parts) >= 2 {
			podName := strings.Trim(parts, "\"'")
			return podName
		}
	}
	return ""
}

/**
 * @description: 判断两个pod item 是否相等. (格式化. 去掉空格和换行去匹配string.equal .)
 * @param {string} item1
 * @param {string} item2
 * @return {*}
 */
func isSamePodItem(item1 string, item2 string) bool {
	item1 = earth.DeleteNewLine(item1)
	item1 = earth.DeleteSpaceSymbol(item1)

	item2 = earth.DeleteNewLine(item2)
	item2 = earth.DeleteSpaceSymbol(item2)

	return item1 == item2
}

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
// @deprecated: This method will be removed in future releases
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
	source = earth.DeleteNewLine(source)
	source = earth.DeleteSpaceSymbol(source)
	if strings.HasPrefix(source, "pod") {
		return true
	}
	if strings.Contains(source, "subspecs") {
		return true
	}
	if strings.Contains(source, "configurations") {
		return true
	}
	if strings.Contains(source, "platform") {
		return true
	}
	if strings.Contains(source, "target") {
		return true
	}
	if strings.Contains(source, "source") {
		return true
	}
	if strings.Contains(source, "path") {
		return true
	}
	if strings.Contains(source, "abstract_target") {
		return true
	}
	if strings.Contains(source, "post_install") {
		return true
	}
	if strings.Contains(source, "binary") {
		return true
	}
	if (strings.HasPrefix(source, "'") && strings.HasSuffix(source, "'")) ||
		(strings.HasPrefix(source, "\"") && strings.HasSuffix(source, "\"")) {
		// 版本号去掉
		return false
	}
	return false
}

// 字符串全部都是空格?
func isAllWhiteSpace(source string) bool {
	source = earth.DeleteSpaceSymbol(source)
	return source == ""
}

// 字符串是注释?
func isAnnoation(source string) bool {
	source = earth.DeleteSpaceSymbol(source)
	return strings.HasPrefix(source, "#")
}

// 字符串是 "target do, end, use_frameworks" 等Podfile中无关的关键字?
func isTargetDoEnd(source string) bool {
	clearedStr := earth.DeleteSpaceSymbol(source)
	if strings.HasPrefix(clearedStr, "#") {
		return true
	}
	if strings.HasPrefix(clearedStr, "if") {
		return true
	}
	if strings.HasPrefix(clearedStr, "target") &&
		strings.HasSuffix(clearedStr, "do") {
		return true
	}
	if clearedStr == "end" {
		return true
	}
	if strings.Contains(clearedStr, "use_frameworks") {
		return true
	}
	if strings.Contains(clearedStr, "source") {
		return true
	}
	if strings.Contains(clearedStr, "platform") {
		return true
	}
	if strings.Contains(clearedStr, "post_install") {
		return true
	}
	if strings.Contains(clearedStr, "config.") {
		return true
	}
	if strings.Contains(clearedStr, "installer.") {
		return true
	}
	if strings.Contains(clearedStr, "target.") {
		return true
	}
	if strings.Contains(clearedStr, "inherit") {
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
	// fmt.Println("🐷搜索" + value + "\n")
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
				if firstWordIsPod(v) {
					// fmt.Println("🐷跳出" + resultString + "\n")
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
