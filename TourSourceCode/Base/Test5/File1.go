package main

import (
	"fmt"
	"os"
)

func WriteFile (path string) {
	f, err := os.Create(path)
	if err != nil{
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile (path string) {
	f, err := os.Open(path)
	if err != nil {
			fmt.Println("err = ", err)
			return
	}

	defer f.Close()
	buf := make([]byte, 1024*2)
	n, err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]) )
}

func main () {
	path := "./demo.text"
	//WriteFile(path)
	ReadFile(path)
}
