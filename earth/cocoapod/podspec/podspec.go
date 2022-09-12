package podspec

import (
	"fmt"
	"goPlay/earth"
	"strings"
)

// podspec 内容
func GetPodSpecContent() string {
	var files []string
	files, _ = earth.GetAllFilePaths(".", files)
	// fmt.Printf("%q \n", files)
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		if strings.Contains(fileName, ".podspec") {
			fileContent := earth.ReadFileFrom(fileName)
			fmt.Printf("\n\n")
			fmt.Print(fileContent)
			fmt.Printf("\n\n")
			return fileContent
		}
	}
	return ""
}

// get版本号
func GetVersion() string {
	source := GetPodSpecContent()
	keyLine := earth.FindFirstChoosenLineString(source, "s.version")
	versionString := strings.Split(keyLine, "=")[1]
	versionString = earth.DeleteQuoteSymbol(versionString)
	versionString = earth.DeleteSpaceSymbol(versionString)
	fmt.Printf("\nget spec version: %q\n\n", versionString)
	return versionString
}
