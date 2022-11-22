package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		FileInfo：文件信息
			interface
				Name()，文件名
				Size()，文件大小，字节为单位
				IsDir()，是否是目录
				ModTime()，修改时间
				Mode()，权限

	*/
	fileInfo, err := os.Stat("E:\\gopath\\src\\Golang-100-Days\\Day16-20(Go语言基础进阶)\\code\\day16\\ab.txt")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Printf("%T\n", fileInfo) // *os.fileStat
	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())
	//是否是目录
	fmt.Println(fileInfo.IsDir()) //IsDirectory false
	//修改时间
	fmt.Println(fileInfo.ModTime())
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-
}
