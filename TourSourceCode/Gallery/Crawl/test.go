package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)
var waitGroup = new(sync.WaitGroup)
func main() {
	//创建多个协程，同时下载多个图片
	os.MkdirAll("pic2016", 0666)
	now := time.Now()

	for i :=1; i<24; i++ {
		//计数器+1
		waitGroup.Add(1)
		go download(i)
	}

	//等待所有协程操作完成
	waitGroup.Wait()
	fmt.Printf("下载总时间:%v\n", time.Now().Sub(now))
	//url := "https://pic.netbian.com/uploads/allimg/211227/235428-1640620468b453.jpg"
	//res, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//
	//fiel, _ := os.Create("img.jpg")
	//defer fiel.Close()
	//
	//io.Copy(fiel, res.Body)
	//ReadFile2("./tt.txt")
	//fileObj, _ := os.Open("./tt.txt")
	//defer fileObj.Close()
	//
	//Contents, _ := ioutil.ReadAll(fileObj)
	//fmt.Println(string(Contents))
	//
	//ioutil.WriteFile("./t3.txt", Contents, 0666)
	//
	//
	//fileObj1, _ := os.OpenFile("./tt.txt", os.O_RDWR|os.O_CREATE, 0666)
	//defer fileObj1.Close()
	//
	//Rd := bufio.NewReader(fileObj1)
	//cont, _ := Rd.ReadSlice('#')
	//fmt.Println(string(cont))
	//
	//Wr := bufio.NewWriter(fileObj1)
	//Wr.WriteString("WriteString writes a ## string.")
	//Wr.Flush()
}


func download(i int ){
	url := fmt.Sprintf("http://pic2016.ytqmx.com:82/2016/0919/41/%d.jpg", i)
	fmt.Printf("开始下载:%s\n", url)
	res,err := http.Get(url)
	if err != nil || res.StatusCode != 200{
		fmt.Printf("下载失败:%s", res.Request.URL)
	}
	fmt.Printf("开始读取文件内容,url=%s\n", url)
	data ,err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Printf("读取数据失败")
	}

	ioutil.WriteFile(fmt.Sprintf("/Golang/Crawl/pic2016/1_%d.jpg", i), data, 0644)

	//if failed, sudo chmod 777 pic2016/

	//计数器-1
	waitGroup.Done()
}

func ReadFile2(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0400)
	fileinfo, _ := file.Stat()
	size := fileinfo.Size() //文件大小，单位是字节，int64
	var read_buffer = make([]byte, size) //直接一步到位
	var length int64 = 0 //标记已经读取了多少字节的内容
	for ; length < size; { //循环读取文件内容
		n, _ := file.Read(read_buffer)
		length += int64(n)
	}
	fmt.Println(string(read_buffer))
}


func ReadFile3(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0400)
	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		//line, err := reader.ReadBytes('\n') //line 是一个[]byte
		line, err := reader.ReadString('\n') //line 是一个 string
		//ReadBytes 与 ReadString 两种方法都可以达到目的
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		fmt.Printf("第%d行：%s", i, line)
		if err == io.EOF { //读到文件结尾就退出循环
			break
		}
	}
}
