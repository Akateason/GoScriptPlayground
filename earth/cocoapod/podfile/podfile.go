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

// Podfile内容
func FetchContent() string {
	fileName := GetPodfileFileName()
	return earth.ReadFileFrom(fileName)
}

// 解析Podfile. 分组
// 1.
// 忽略 纯\n
// 忽略 target do ... end  嵌套
// 忽略 #注释
// 2.
// 根据pod内容分组
func Analysis() []string {
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
	for _, value := range resultList {
		fmt.Println(value + "🐷🐶")
	}

	return resultList
}

// 3.
// 导出新Podfile
func ExportNewPodfile() string {
	// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	oldPodfile := FetchContent()
	resultList := Analysis()
	for _, value := range resultList {
		oldStr := findSourceLineWith(value, oldPodfile)
		// fmt.Println("搜索" + value + "\n")
		// fmt.Println("得出" + oldStr + "\n")
		if len(oldStr) > 0 {
			oldPodfile = strings.Replace(oldPodfile, oldStr, value, 1)
		}
	}
	// fmt.Println(" 🐲🐲🐲🐲🐲🐲🐲 ")
	return oldPodfile
}

// ///////PRIVATE/////////
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
		resultString = ""
	}
	return resultString
}
