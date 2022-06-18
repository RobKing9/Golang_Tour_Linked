package main

import (
	"fmt"
	"os"
	"time"
)

func main () {
	fmt.Println(time.Now())
	fmt.Println(os.Hostname())
	fmt.Println(os.Getwd())
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getpid())
	fmt.Fprintf(os.Stdout, "%s\n", "hellogo")
	fmt.Println(time.Now())
}
