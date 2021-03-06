/**
@Title
@Description
@Author:king 2022/2/9 下午11:28
@Update      2022/2/9 下午11:28
**/
package file

import (
	"fmt"
	"io"
	"os"
	"path"
)



func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetFileSuffix(fileName string) string {
	filesuffix := path.Ext(fileName)

	return filesuffix
}



func WriteFile(filename string, content string) error {
	var f *os.File
	var err1 error
	if FileIsExist(filename) { //如果文件存在
		os.Remove(filename)
		//
		//f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		//fmt.Println("文件存在")
	} else {
		//f, err1 = os.Create(filename) //创建文件
		//fmt.Println("文件不存在")
	}

	f, err1 = os.Create(filename) //创建文件
	if err1 != nil {
		return err1
	}
	//check(err1)
	defer f.Close()

	n, err1 := io.WriteString(f, content) //写入文件(字符串)
	//check(err1)
	if err1 != nil {
		return err1
	}
	fmt.Printf("写入 %d 个字节n", n)

	return nil
}



func FileIsExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func FolderIsExist(path string, create bool) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) && create {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
		} else {
			return true, nil
		}
	}
	return false, err
}

func CreateFolder(path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}