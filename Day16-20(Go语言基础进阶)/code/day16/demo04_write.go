package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		写出数据：
	*/

	fileName := "Day16-20(Go语言基础进阶)/code/day16/bb.txt"
	//step1：打开文件
	//step2：写出数据
	//step3：关闭文件
	//file,err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//写出数据
	//bs :=[]byte{65,66,67,68,69,70}//A,B,C,D,E,F
	bs := []byte{97, 98, 99, 100} //a,b,c,d
	//n,err := file.Write(bs)
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	Student164(err)
	file.WriteString("\n")

	//直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	Student164(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	Student164(err)

}

func Student164(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
