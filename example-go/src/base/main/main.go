package main

import "fmt"

func main()  {

	fmt.Printf("hello world! \n")

	// for
	for i := 0; i < 10; i++ {
		fmt.Printf("for test %v \n", i)
	}


	// while
	sum := 0
	for sum < 10 {
		fmt.Printf("while test %v \n", sum)
		sum++
	}

	// 类型转换
	var var1 int = 11

	var2 := float64(var1)
	var3 := float32(var1)

	fmt.Printf("type change var1: %T -> %v, var2: %T -> %v, var3: %T -> %v\n", var1, var1, var2, var2, var3, var3)


	// 函数调用



}


/**
	平方根函数
 */
func Sqrt(x float64) float64 {
	if x <= 0 {
		panic("param must not be zero")
	}


}