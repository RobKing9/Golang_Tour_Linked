package write

import (
	"fmt"
	"os"
)

//file.write

func WriteFile1() {
	//创建目录
	//单级
	//os.Mkdir("writefile", 0744)
	//多级
	os.MkdirAll("writefile/w1", 0744)
	fileobj, err := os.OpenFile("./writefile/w1/w1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err.Error())
	}
	defer fileobj.Close()
	str := "你好 世界"
	//fileobj.Write([]byte(str))
	fileobj.WriteString(str)
}