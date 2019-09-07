package main

import (
	"fmt"
)

func printArray(array []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)    // cap = capacity
}

func test1() {
    var array1 = []int{2, 3, 4, 5, 6, 7}

    fmt.Printf("array1: '%v' type=%T\n", array1, array1)
}

func test2()  {
    var array1 = []int{2, 3, 4, 5, 6, 7}
    var i int

    for i=0; i<len(array1); i++ {
        fmt.Printf("%v: %v type=%T\n", i, array1[i], array1[i])
    }
}

func test3() {
	var array1 = []int{2, 3, 4, 5, 6, 7}
    printArray(array1)
}

func test4() {
	var array1 = []int{2, 3, 4, 5, 6, 7}
    printArray(array1)
}

func test5() {
    array := make([]int, 5)
    printArray(array)
}

func test6() {
    array1 := []int{1, 2, 3, 4, 5, 6, 7}
    array2 := array1[:3]
    array3 := array1[4:]

    printArray(array1)
    printArray(array2)
    printArray(array3)
}


func test7() {
    var array1 = []int{}

    array1 = append(array1, 2)
    array1 = append(array1, 3, 4, 5)

    printArray(array1)

}

func test8() {
    array1 := []string{"a", "b", "c", "d", "e"}

    for i, v := range array1 {
		fmt.Printf("%d: %s\n", i, v)
	}
}


func test9() {
    array1 := []string{"a", "b", "c", "d", "e"}

    for i, v := range array1 {
		fmt.Printf("%d: %s\n", i,  v)
	}
}




func main() {
	//test1()
    //test2()
    //test3()
    //test5()
    //test6()
    //test7()
    //test8()
    test9()

}
