package main

import (
	"os"
	"syscall"
	"testing"
)

const GDATA = 1024 * 1024

var strslice = make([]byte, 0, GDATA)

func init() {
	for i := 0; i < GDATA; i++ {
		strslice = append(strslice, byte(i))
	}
}

func TestWrite2RowIO(t *testing.T) {
	fd, err := syscall.Open("./test3.txt", os.O_WRONLY|os.O_SYNC|syscall.O_CREAT|syscall.O_EXCL, 0666)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)
	syscall.Write(fd, strslice)
}

func TestWriteLinuxIO(t *testing.T) {
	fd, err := syscall.Open("./test4.txt", os.O_WRONLY|syscall.O_CREAT|syscall.O_EXCL, 0666)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)
	syscall.Write(fd, strslice)
}
