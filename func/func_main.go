package main

import (
	"fmt"
)

func test1() {
	fmt.Println("Testing 1")
}

func test2() (a int, b int, c int) {
	return 1, 2, 3
}

func test3(num1, num2, num3 int) int {
	return num1 + num2 + num3
}


func main() {
	//fmt.Println("Hello Go!")
	test1()

	a, b, c := test2()
	var sum = test3(2, 3, 4)


	fmt.Println("result: a=", a, " b=", b, " c=", c)
	fmt.Println("result: sum=", (a + b + c))
	fmt.Println("result: sum=", sum)
}
