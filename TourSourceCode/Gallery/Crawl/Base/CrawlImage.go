package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)
var (

	reImg = "(https?:[^:<>\"]*\\/)([^:<>\"]*)(\\.((png!thumbnail)|(png)|(jpg)|(webp)))"
)

func CrawlPageStr (url string) (PageStr string) {
	res,err := http.Get(url)
	HandleError(err, "http.Get url")
	defer res.Body.Close()
	//读取页面内容
	PageByte, err := ioutil.ReadAll(res.Body)
	HandleError(err, "ioutil.ReadAll")
	PageStr = string(PageByte)
	return PageStr
}

func HandleError (err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func CrawlImg (url string) {
	PageStr := CrawlPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(PageStr, -1)
	for _, result := range results {
		fmt.Println("Img = ", result[0])
	}
}



func main () {
	CrawlImg("https://umei.cc/bizhitupian/meinvbizhi/xingganmeinv.htm")
}
