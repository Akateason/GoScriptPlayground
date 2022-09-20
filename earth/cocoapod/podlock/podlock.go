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
	fmt.Println("解析开始, 以🐷🐶结尾")
	lockContent := FetchContent()
	sourceList := strings.Split(lockContent, "\n\n")
	// 打印解析结果
	for _, value := range sourceList {
		fmt.Println(value + "🐷🐶")
	}
	return sourceList
}
