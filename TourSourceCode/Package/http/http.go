package main

import (
	"fmt"
	"net/http"
)

func main() {
	//get方法请求网页
	resp, err := http.Get("https://www.bilibili.com")
	if err != nil {
		fmt.Println("http.get failed:", err.Error())
	}
	defer resp.Body.Close()

	//网页头部
	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
	//网页主体
	//body := resp.Body
	//fmt.Println(body)
}
