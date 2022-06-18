package main

import "fmt"

func main () {
	m1 := make (map [int] string, 2)
	m1[1] = "wen"
	m1[2] = "zeng"
	m1[3] = "xiang"
	fmt.Println("m1 = ", m1)
	fmt.Println("len = ", len(m1))
	delete(m1, 1)
	fmt.Println("m1 = ", m1)

	m2 := map[int]string{1: "zeng", 2: "xiang", 3: "wen"}
	fmt.Println("m2 = ", m2)
	m2[1] = "dog"
	fmt.Println("m2 = ", m2)
	m2[9] = "cat"
	fmt.Println("m2 = ", m2)
	for key, value := range m2 {
		fmt.Printf("%d =======> %s\n", key, value)
	}


}
