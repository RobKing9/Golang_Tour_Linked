package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	rePhone = "1[3456789]\\d\\s?\\d{4}\\s?\\d{4}"
	//reIDCard = "[123456789]\\d{5} ( (19\\d{2}) | (20[01]\\d) ) ( (0[1-9]) | (1[012]) ) ( (0[1-9]) | ([12]\\d) | (3[01]) ) \\d{3}[\\dXx] "
	//reIDCard = "/^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$/"
)

func HandleError (err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func CrawPhone (url string) {
	//获取数据
	res, err := http.Get(url)
	HandleError(err,"http.Get url")
	defer res.Body.Close()
	//读取页面内容
	PageByte,err := ioutil.ReadAll(res.Body)
	HandleError(err,"ioutil.ReadAll")
	PageStr := string(PageByte)
	//筛选内容
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(PageStr, -1)
	for _, result := range results {
		fmt.Println("Tel = ",result[0])
	}
}

//func CrawlIDCard (url string) {
//	//获取数据
//	res, err := http.Get(url)
//	HandleError(err,"http.Get url")
//	defer res.Body.Close()
//	//读取页面内容
//	PageByte,err := ioutil.ReadAll(res.Body)
//	HandleError(err,"ioutil.ReadAll")
//	PageStr := string(PageByte)
//	fmt.Println(PageStr)
//	//筛选内容
//	re := regexp.MustCompile(reIDCard)
//	results := re.FindAllStringSubmatch(PageStr, -1)
//	fmt.Println(results)
//	for _, result := range results {
//		fmt.Println("IDCard = ",result[0])
//	}
//
//
//}



func main () {
	CrawPhone("https://www.zhaohaowang.com/")
	//CrawlIDCard("https://baijiahao.baidu.com/s?id=1610648123471451348&wfr=spider&for=pc")
}
