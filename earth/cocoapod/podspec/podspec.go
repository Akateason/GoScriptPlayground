package podspec

import (
	"fmt"
	"goPlay/earth"
	"strings"
)

// 获取对应podSpec的文件名
func getSpecFileName() string {
	var files []string
	files, _ = earth.GetAllFilePaths(".", files)
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		if strings.Contains(fileName, ".podspec") {
			return fileName
		}
	}
	return ""
}

// podspec 内容
func GetPodSpecContent() string {
	fileName := getSpecFileName()
	return earth.ReadFileFrom(fileName)
}

// get spec 版本号
func GetVersion() string {
	source := GetPodSpecContent()
	keyLine := earth.FindFirstChoosenLineString(source, "s.version")
	versionString := strings.Split(keyLine, "=")[1]
	versionString = earth.DeleteQuoteSymbol(versionString)
	versionString = earth.DeleteSpaceSymbol(versionString)
	fmt.Printf("get spec version: %q\n\n", versionString)
	return versionString
}

// 更新版本号
// index: 版本的第几位-> 0,1,2,  0是最大版本, 2是最小版本, 默认为2
func UpdateVersion(index int) {
	willUpdateVersionIndex := index

	oldVersion := GetVersion()

	vItemList := strings.Split(oldVersion, ".")
	item := vItemList[willUpdateVersionIndex]
	intItem := earth.Str2Int(item)
	intItem++
	vItemList[willUpdateVersionIndex] = earth.Int2Str(intItem)

	newVersion := strings.Join(vItemList, ".")
	result := "\ts.version\t= '" + newVersion + "'"
	fmt.Printf("success 🚀🚀🚀 new Version is: %q \n\n", newVersion)

	podspecSource := GetPodSpecContent()
	keyLine := earth.FindFirstChoosenLineString(podspecSource, "s.version")

	podspecSource = strings.Replace(podspecSource, keyLine, result, -1)
	earth.WriteStringToFileFrom(getSpecFileName(), podspecSource)
}
