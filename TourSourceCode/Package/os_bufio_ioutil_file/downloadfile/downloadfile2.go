package downloadfile

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile2(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get failed, err:", err.Error())
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		fmt.Println("open file failed, err:", err.Error())
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("copy failed, err:", err.Error())
	}
}
