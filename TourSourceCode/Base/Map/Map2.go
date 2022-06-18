package main

import (
	"fmt"
	"time"
)

func test (m map[int]string) {
	delete(m, 2)
}

func main () {
	var m map[int]string{1: "go", 2: "java", 3: "python"}
	test(m)
	fmt.Println("m = ", m)
}
