package main

import (
	"fmt"
)

func test_int() {
    // Boolean, Numerics, String
    // Testing Typed Variable
    var int1 int                // will use 0
    var int2 int = 10
    var int3 = 20
    var int4 = -20
    //var int5 uint8=1000         // this will cause overflow
    //var int5 uint8=-1         // this will cause overflow
    var int5 uint8=255         // this will cause overflow
    var int6 int64=9223372036854775807


	fmt.Printf("int 1: %v type=%T\n", int1, int1)
    fmt.Printf("int 2: %v type=%T\n", int2, int2)
    fmt.Printf("int 3: %v type=%T\n", int3, int3)
    fmt.Printf("int 4: %v type=%T\n", int4, int4)
    fmt.Printf("int 5: %v type=%T\n", int5, int5)
    fmt.Printf("int 6: %v type=%T\n", int6, int6)

}

func test_float() {
    // Boolean, Numerics, String
    // Testing Typed Variable
    var float1 float32 = 0.3
    var float2 = 0.3
    var float3 = 0.3141567
    var complex1 = 1 + 2i
    var complex2 complex64

    complex2 = 2 + 3i


    fmt.Printf("float 1: %v type=%T\n", float1, float1)
    fmt.Printf("float 2: %v type=%T\n", float2, float2)
    fmt.Printf("float 3: %v type=%T\n", float3, float3)
    fmt.Printf("complex 1: %v type=%T\n", complex1, complex1)
    fmt.Printf("complex 2: %v type=%T\n", complex2, complex2)
}

func test_bool() {
    // Boolean, Numerics, String
    // Testing Typed Variable
    var flag1 bool = true
    var flag2 bool = false
    var flag3 bool

    flag3 = true

    fmt.Printf("bool 1: %v type=%T\n", flag1, flag1)
    fmt.Printf("bool 2: %v type=%T\n", flag2, flag2)
    fmt.Printf("bool 3: %v type=%T\n", flag3, flag3)
}


func test_string() {
    // Boolean, Numerics, String
    // Testing Typed Variable
    var str1 string
    var str2 string = "abcde"
    var str3 string

    str3 = "xyz"

    fmt.Printf("str 1: '%v' type=%T\n", str1, str1)
    fmt.Printf("str 2: '%v' type=%T\n", str2, str2)
    fmt.Printf("str 3: '%v' type=%T\n", str3, str3)

}

func test_int_array() {
    // Boolean, Numerics, String
    // Testing Typed Variable
    var array1 = []int{2, 3, 4, 5, 6, 7}

    fmt.Printf("array1: '%v' type=%T\n", array1, array1)

}

func short_var() {
	//https://golang.org/ref/spec#Short_variable_declarations
	// no need the var
	a := 1
	b := 0.5

	fmt.Printf("a=%v(%T) b=%v(%T)\n", a, a, b, b)
}


func main() {
	//fmt.Println("Hello Go!")
	test_int();
    test_float();
    test_bool();
    test_string();
	test_int_array();
	short_var();

    //
	// a, b, c := test2()
	// fmt.Println("result: a=", a, " b=", b, " c=", c)
	// fmt.Println("result: sum=", (a + b + c))
}
