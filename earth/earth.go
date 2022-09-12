package earth

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

// 调用命令行
func UseCommandLine(cmd string) error {
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

// readFileFrom 使用ioutil.ReadFile 直接从文件读取到 []byte中
func ReadFileFrom(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("读取文件失败:%#v", err)
		return ""
	}
	return string(f)
}

// isFileExists 判断所给路径文件/文件夹是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// ifNoFileToCreate 文件不存在就创建文件
func IfNoFileToCreate(fileName string) (file *os.File) {
	var f *os.File
	var err error
	if !IsFileExists(fileName) {
		f, err = os.Create(fileName)
		if err != nil {
			return
		}
		log.Printf("IfNoFileToCreate 函数成功创建文件:%s", fileName)
		defer f.Close()
	}
	return f
}

// writeStringToFileFrom 通过 ioutil.WriteFile 写入文件
func WriteStringToFileFrom(fileName string, writeInfo string) {
	_ = IfNoFileToCreate(fileName)
	info := []byte(writeInfo)
	if err := ioutil.WriteFile(fileName, info, 0666); err != nil {
		log.Printf("WriteStringToFileFrom %q 写入文件失败:%+v", fileName, err)
		return
	}
	log.Printf("WriteStringToFileFrom %q 写入文件成功", fileName)
}

// 获取当前项目根目录下所有文件
func GetAllFilePaths(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// 通过 keyword 找对应行
func FindFirstChoosenLineString(source string, keyword string) string {
	var list = strings.Split(source, "\n")
	for i := 0; i < len(list); i++ {
		var tmpStr = list[i]
		if strings.Contains(tmpStr, keyword) {
			return tmpStr
		}
	}
	return ""
}

// 删除单双引号
func DeleteQuoteSymbol(source string) string {
	var tmp = strings.Replace(source, "'", "", -1)
	tmp = strings.Replace(tmp, "\"", "", -1)
	return tmp
}

// 删除空格
func DeleteSpaceSymbol(source string) string {
	return strings.Replace(source, " ", "", -1)
}

// int to string
func Int2Str(num int) string {
	return strconv.Itoa(num)
}

// string to int
func Str2Int(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
