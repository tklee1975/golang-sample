package main

/**
 * # How to Run
 * Run all test
 * >go test mypack_test.go
 *
 * Run all with detail
 * >go test basic/utest/mypack_test.go  -v
 *
 * Run specific test
 * >go test basic/utest/mypack_test.go  -v -run Simple
 *
 * Benchmark test
 * >go test basic/utest/mypack_test.go -v -bench=.
 *
 * # Reference
 * - https://golang.org/pkg/testing/
 * - https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file
 * - https://openhome.cc/Gossip/Go/Testing.html
 */

import (
    "testing"
    "fmt"
)

func getArray(size int) []int {
    var array []int

    for i:=0; i<size; i++ {
        array = append(array, (i+1))
    }

    return array;
}

// Item the type of the Set
type Item generic.Type

// Generic Reverse
func Reverse (type  Element) (s []Element) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}

func Test_ReverseInt(t *testing.T) {
    var array []int

    array = getArray(10)

    fmt.Printf("%v\n", array);
}
