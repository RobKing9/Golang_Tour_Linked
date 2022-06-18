package write

import (
	"bufio"
	"fmt"
	"os"
)

//bufio.NewWriter

func WriteFile2() {
	fileobj, err := os.OpenFile("./writefile/w1/w2.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err.Error())
	}
	defer fileobj.Close()

	writer := bufio.NewWriter(fileobj)
	i := 0
	for {
		i++
		writer.WriteString("你好 世界\n") //写入缓存
		if i == 10 {
			break
		}
	}
	writer.Flush() //将缓存中的内容写入文件
}