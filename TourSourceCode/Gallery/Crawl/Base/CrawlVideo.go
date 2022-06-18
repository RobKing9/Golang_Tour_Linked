package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reVideo = "(https?:[^:<>\"]*\\/)([^:<>\"]*)(\\.((png!thumbnail)|(png)|(jpg)|(webp)))"
)

func CrawlVideo (url string) {
	//获取数据
	res, err := http.Get(url)
	HandleError(err,"http.Get url")
	defer res.Body.Close()
	//读取页面内容
	PageByte,err := ioutil.ReadAll(res.Body)
	HandleError(err,"ioutil.ReadAll")
	PageStr := string(PageByte)
	//筛选内容
	re := regexp.MustCompile(reVideo)
	results := re.FindAllStringSubmatch(PageStr, -1)
	for _, result := range results {
		fmt.Println("Video = ",result[0])
	}
}

func HandleError (err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main () {
	CrawlVideo("https://www.bilibili.com/video/BV1PP4y1j75j?spm_id_from=333.999.0.0")
}

