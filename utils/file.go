package md5

import (
	"io/ioutil"
	"os"
)


func scanFile(level int, fileDir string, relative string, fileCache [] string) {

	pathSeparator := string(os.PathSeparator)

	//fmt.Println("################### fileDir :", fileDir)
	//fmt.Println("################### relative :", relative)
	files, _ := ioutil.ReadDir(fileDir)
	if files == nil {
		//fmt.Println("################### searchPathFiles files is nil")
		return
	}
	for _, onefile := range files {
		if onefile.IsDir() {
			//fmt.Println(fileDir + pathSeparator + onefile.Name())
			//fmt.Println(pathSeparator + onefile.Name())
			scanFile(level+1, fileDir+pathSeparator+onefile.Name(), relative+pathSeparator+onefile.Name(), fileCache)
		} else {
			var findFilePath = relative + pathSeparator + onefile.Name()
			//var findFilePath = onefile.Name()
			fileCache = append(fileCache, findFilePath)
			//fmt.Println("-----", findFilePath )
		}
	}
}

func ScanFile(fileDir string) [] string {
	var files = make([] string, 0, 1)
	scanFile(1, fileDir, "", files)

	//fmt.Println("-----------------" )
	//fmt.Println("-----共找到文件数量：", len(files) )
	//fmt.Println("-----------------" )
	//for k := range files {
	//	fmt.Println("找到了文件 ： ", k)
	//}

	return files
}
