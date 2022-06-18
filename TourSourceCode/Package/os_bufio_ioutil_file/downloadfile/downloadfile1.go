package downloadfile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadFile1(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get failed, err:", err.Error())
	}
	defer resp.Body.Close()

	//创建目录
	os.MkdirAll("./image", 0744)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read failed, err:", err.Error())
	}
	ioutil.WriteFile(filename, buf, 0666)
}
