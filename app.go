package main

import (
	"fmt"
	"go-helper/file"
	"os"
)

func main() {
	osPath , _ := os.Getwd()
	exist := file.FileIsExist(osPath+"/README.md")
	if exist {
		fmt.Println("file exist")
	}else {

		fmt.Println("file not exist")
	}

	exist2,_ := file.FolderIsExist(osPath+"/file11", false)
	if exist2 {
		fmt.Println("Folder Is Exist")
	}else {

		fmt.Println("Folder Is Not Exist")
	}
}
