package main

import (
	"encoding/json"
	"fmt"
)

func main () {
	m := make(map[string]interface{}, 4)
	m["company"] = "incast"
	m["subject"] = []string{"Go", "C++", "Python"}
	m["isok"] = true
	m["price"] = 333.333
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
}

