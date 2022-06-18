package main

import "fmt"

type Student struct{
	name string
	age int
}

func main () {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student{"fumin", 17}

	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int, 内容为%s\n", index, value)
		}else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		}else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student, 内容为name = %s, age = %d\n", index, value.name, value.age)
		}

}

