package write

import (
	"fmt"
	"io/ioutil"
)

func WriteFile3() {
	str := "你好 世界"
	err := ioutil.WriteFile("./writefile/w1/w3.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err.Error())
	}
}
