package read

import (
	"fmt"
	"io/ioutil"
)

//ioutil 一次性读取文件

func ReadFile3() {
	buf, err := ioutil.ReadFile("./hllo.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err.Error())
	}
	fmt.Println(string(buf))
}

