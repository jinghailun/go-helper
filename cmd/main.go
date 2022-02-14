package main

import (
	"fmt"
	utilsHelper "github.com/jinghailun/go-helper/utils"
)
func main() {
	fileList := utilsHelper.ScanFile("F:/project/ht_project/config_cache")

	for _, fileName := range fileList {
		fmt.Println("file : ", fileName)
	}
}
