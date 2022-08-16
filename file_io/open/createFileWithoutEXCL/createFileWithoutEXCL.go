package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func createFileWithoutEXCL(path string) error {

	// 判断文件是否存在
	if _, exist := os.Stat(path); exist != nil {
		fmt.Printf("file %s not exist\n", path)
	} else {
		panic("file already exists")
	}

	// 创建一个同名文件
	go func(path string) {
		fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_CREAT, 0666)
		if err != nil {
			panic(err.Error())
		}
		syscall.Write(fd, []byte("abc"))
	}(path)

	time.Sleep(1 * time.Second)

	fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_CREAT, 0666)
	if err != nil {
		return err
	}

	syscall.Write(fd, []byte("def\n"))

	return nil
}

func main() {
	createFileWithoutEXCL("./test.file")
}
