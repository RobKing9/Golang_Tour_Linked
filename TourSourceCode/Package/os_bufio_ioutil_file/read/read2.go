package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//bufio按行读取文件

func ReadFile2() {
	//打开文件
	fileobj, err := os.OpenFile("./hello.txt", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println("open file failed, err:", err.Error())
	}
	//关闭文件
	defer fileobj.Close()
	//读取文件
	reader := bufio.NewReader(fileobj)
	for {
		str, err := reader.ReadString('\n')
		if err==io.EOF {
			if len(str) != 0 {
				fmt.Println(str)
			}
			fmt.Println("read over!")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err.Error())
		}
		//fmt.Printf("第%d行:%s\n", i, str)
		fmt.Println(str)
	}
}
