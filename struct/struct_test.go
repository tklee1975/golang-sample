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

// Struct definition

// Rect Structure
type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r *rect) perimeter() int {
    return (r.width + r.height) * 2
}

func (r *rect) info() string {
    return fmt.Sprintf("%v x %v", r.width, r.height);
}

// Person Structure
type person struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r *rect) perimeter() int {
    return (r.width + r.height) * 2
}

func (r *rect) info() string {
    return fmt.Sprintf("%v x %v", r.width, r.height);
}



func TestRect(t *testing.T) {
    r := rect{ width:10, height:6 }

    fmt.Printf("size      : %v\n", r.info())
    fmt.Printf("area      : %v\n", r.area())
    fmt.Printf("perimeter : %v\n", r.perimeter())
    fmt.Printf("string    : %v\n", r)
}
