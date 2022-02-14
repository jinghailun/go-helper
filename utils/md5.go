package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

const filechunk = 8192


func Md5SumBigFile(filePath string) (value []byte, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()
	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	value = hash.Sum(nil)

	//fmt.Printf("%s checksum is %x\n", file.Name(), hash.Sum(nil))
	return
}

func Md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5File(file string) (value []byte, err error) {
	value, err = Md5SumBigFile(file)
	return
}