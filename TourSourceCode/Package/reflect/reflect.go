package main

import (
	"fmt"
	"reflect"
)

func main() {
	//reflect.TypeOf()获取对象类型信息
	//reflect.ValueOf() 获取对象值
	var a int64 = 100
	fmt.Printf("Type:%v, Value:%v", reflect.TypeOf(a), reflect.ValueOf(a))
}
