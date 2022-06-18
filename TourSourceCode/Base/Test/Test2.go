package main

import "fmt"

func main() {
	a := 10
	str := "hello"
	func () {
		fmt.Printf("a = %d,str = %s \n " , a , str)
	}()

	func (i,j int) {
		fmt.Printf("i= %d , j = %d \n" , i , j)
	} (10 , 20)

	sum := func ( x ,y int ) {
		fmt.Println( x+y )
	}
	sum (34 , 45)
}
