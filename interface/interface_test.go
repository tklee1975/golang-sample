package interface_test

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
 *
 * - https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c
 */

import (
    "testing"
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
   width, height float64
}

type circle struct {
   radius float64
}

type square struct {
   size float64
}


// function implements the interface
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2 * r.width + 2 * r.height
}

// function implements the interface
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

//
func measure(g geometry) {
    fmt.Printf("Info: %v type:%T\n", g, g)
    fmt.Println("Area: ", g.area())
    fmt.Println("Perim: ", g.perim())
}


func Test_Measure(t *testing.T) {
    r := rect{ width:3, height:4 }
    c := circle{ radius:5 }
    s := square{ size:6 }

    measure(r)
    measure(c)
    // measure(s)      // error: square does not implement geometry (missing area method)
}
