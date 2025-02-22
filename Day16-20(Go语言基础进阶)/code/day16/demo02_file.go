package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	/*
		文件操作：
		1.路径：
			相对路径：relative
				ab.txt
				相对于当前工程
			绝对路径：absolute
				/Users/ruby/Documents/pro/a/aa.txt

			.当前目录
			..上一层
		2.创建文件夹，如果文件夹存在，创建失败
			os.MkDir()，创建一层
			os.MkDirAll()，可以创建多层

		3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
			os.Create()，创建文件

		4.打开文件：让当前的程序，和指定的文件之间建立一个连接
			os.Open(filename)
			os.OpenFile(filename,mode,perm)

		5.关闭文件：程序和文件之间的链接断开。
			file.Close()

		5.删除文件或目录：慎用，慎用，再慎用
			os.Remove()，删除文件和空目录
			os.RemoveAll()，删除所有
	*/
	//1.路径
	fileName1 := "Day16-20(Go语言基础进阶)/code/day16/ab.txt"
	//fileName2 := "E:\\gopath\\src\\Golang-100-Days\\Day16-20(Go语言基础进阶)\\code\\day16\\bb.txt"
	//fmt.Println(filepath.IsAbs(fileName1)) //true
	//fmt.Println(filepath.IsAbs(fileName2)) //false
	//fmt.Println(filepath.Abs(fileName1))
	//fmt.Println(filepath.Abs(fileName2)) // E:\gopath\src\Golang-100-Days\bb.txt <nil>
	//
	//fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	//2.创建目录
	//err := os.Mkdir("Day16-20(Go语言基础进阶)/code/day16/test", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("文件夹创建成功。。")
	//err :=os.MkdirAll("/Users/ruby/Documents/pro/a/cc/dd/ee",os.ModePerm)
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("多层文件夹创建成功")

	//3.创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	//file1, err := os.Create("Day16-20(Go语言基础进阶)/code/day16/test/test.txt")
	//if err != nil {
	//	fmt.Println("err：", err)
	//	return
	//}
	//fmt.Println(file1)
	//fileName3 := "Day16-20(Go语言基础进阶)/code/day16/test/test2.txt"
	//file2, err := os.Create(fileName3) //创建相对路径的文件，是以当前工程为参照的
	//if err != nil {
	//	fmt.Println("err :", err)
	//	return
	//}
	//fmt.Println(file2)

	//4.打开文件：
	//file3, err := os.Open(fileName1) //只读的
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file3)
	/*
		第一个参数：文件名称
		第二个参数：文件的打开方式
			const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
			O_RDONLY int = syscall.O_RDONLY // open the file read-only.
			O_WRONLY int = syscall.O_WRONLY // open the file write-only.
			O_RDWR   int = syscall.O_RDWR   // open the file read-write.
			// The remaining values may be or'ed in to control behavior.
			O_APPEND int = syscall.O_APPEND // append data to the file when writing.
			O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
			O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
			O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
			O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
		)
		第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	*/
	file4, err := os.Open(fileName1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)
	all, err := ioutil.ReadAll(file4)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return
	}
	fmt.Printf(string(all))
	//5关闭文件，
	defer func(file4 *os.File) {
		err := file4.Close()
		if err != nil {
			fmt.Println("err:", err)
		}
	}(file4)

	//6.删除文件或文件夹：
	//删除文件
	//err :=  os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("删除文件成功。。")
	//删除目录
	//err := os.RemoveAll("/Users/ruby/Documents/pro/a/cc")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("删除目录成功。。")
}

func WriteTxt(file *os.File) {
	// 打开文件路径 "D:/helloworld.txt"
	// 只写方式打开 os.O_WRONLY
	// os.FileMode(0600) 文件权限：windows系统权限失效。
	// 关闭文件
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//
	//	}
	//}(file)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// 写入一行
			_, err := file.WriteString(strconv.Itoa(j) + " ")
			if err != nil {
				log.Fatal(err)
			}
		}
		file.WriteString("\r\n")
	}
}

func ReadTxt() {
	file, err := os.Open("Day16-20(Go语言基础进阶)/code/day16/ab.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf(line)
	}
}
