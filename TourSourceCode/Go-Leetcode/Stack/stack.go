package main

import (
	"fmt"
	"sync"
)

//栈的数据结构

type Stack struct {
	//数组存放值
	array []string
	//大小
	size int
	lock sync.Mutex //为了并发安全
}

//入栈

func (s *Stack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.array = append(s.array, v)

	s.size = s.size + 1
}

//出栈

func (s *Stack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	//栈中元素为空
	if s.size == 0 {
		panic("empty")
	}
	//取栈顶元素
	top := s.array[s.size-1]
	//创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newarray := make([]string, s.size-1, s.size-1)
	for i := 0; i < s.size-1; i++ {
		newarray[i] = s.array[i]
	}
	s.array = newarray
	s.size = s.size - 1
	return top
}

//去栈顶元素

func (s *Stack) GetTop() string {
	//栈中元素为空
	if s.size == 0 {
		panic("empty")
	}

	top := s.array[s.size-1]
	return top
}

//栈的大小

func (s *Stack) Size() int {
	return s.size
}

func main() {
	//入栈
	fmt.Println("hello")
	/*
		newarray := new(Stack)
		newarray.Push("zxw")
		newarray.Push("robking")
		newarray.Push("fumin")
		fmt.Printf("after push size = %d \n", newarray.Size())
		fmt.Printf("top = %s\n", newarray.GetTop())
		fmt.Printf("pop1 = %s\n", newarray.Pop())
		fmt.Printf("pop2 = %s\n", newarray.Pop())
		fmt.Printf("after pop size = %d\n", newarray.Size())

	*/

}
