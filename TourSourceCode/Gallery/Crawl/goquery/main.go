package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("http.get failed:", err.Error())
	}
	defer req.Body.Close()
	//defer func() { _ = req.Body.Close }()
	dom, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		fmt.Println("goquery.read failed:", err.Error())
	}
	dom.Find("script[id='RENDER_DATA']").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
