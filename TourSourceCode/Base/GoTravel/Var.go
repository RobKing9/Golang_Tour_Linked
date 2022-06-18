package main

import (
	"fmt"
	//"math/cmplx"

	"math"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	//z complex64 = cmplx.Sqrt()
)

func main () {
	fmt.Printf("Type : %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type : %T Value: %v\n", MaxInt, MaxInt)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z int = int(f)
	fmt.Println(x, y, f, z)

}
