package read

import (
	"fmt"
	"io"
	"os"
)

func ReadFile1() {
	//打开文件
	//fileobj, err := os.Open("./hello.txt")
	fileobj, err := os.OpenFile("./hello.txt", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println("open file failed:", err.Error())
	}
	//关闭文件
	defer fileobj.Close()
	//读取文件
	//文件的信息
	fileinfo, err := fileobj.Stat()
	if err != nil {
		fmt.Println("fileinfo failed:", err.Error())
	}
	size := fileinfo.Size()
	fmt.Printf("这个文件大小为%d\n", size)

	var buf = make([]byte, size)
	//循环读取
	//var length int64
	//for ; length<size; {
	//	n,err := fileobj.Read(buf)
	//	if err == io.EOF {
	//		fmt.Println("文件读取完毕！")
	//	}
	//	if err != nil {
	//		fmt.Println("read file failed:", err.Error())
	//	}
	//	length+=int64(n)
	//}
	//fmt.Printf("读取了%d个字节", length)
	//fmt.Println()
	//fmt.Println(string(buf[:length]))
	n, err := fileobj.Read(buf)
	if err == io.EOF {
		fmt.Println("文件读取完毕！")
	}
	if err != nil {
		fmt.Println("read file failed:", err.Error())
	}
	fmt.Printf("读取了%d个字节", n)
	fmt.Println()
	fmt.Println(string(buf[:n]))
}
