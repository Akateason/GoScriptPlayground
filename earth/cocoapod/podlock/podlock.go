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

func fetchPOD() string {
	analysisResult := Analysis()
	for _, v := range analysisResult {
		if v[0:5] == "PODS:" {
			return v
		}
	}
	return ""
}

func CheckPodVersion(podName string) []string {
	podName = strings.ToLower(podName)
	POD := fetchPOD()
	fmt.Println("🦆🦆🦆搜索🦆🦆🦆" + podName + "的版本号")
	sourceList := strings.Split(POD, "\n")
	var resultList []string
	for _, v := range sourceList {

		if v[0:3] == "  -" {
			var lowerV = strings.ToLower(v)
			if strings.Contains(lowerV, podName) {
				resultList = append(resultList, v)
				fmt.Println("🦆" + v)
			}
		}
	}
	return resultList
}
