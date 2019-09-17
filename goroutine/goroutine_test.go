package goroutine_test

//  use
//  go test goroutine_test.go
//  go benchmark goroutine_test.go
// to test
//
// Reference:
//  https://gobyexample.com/goroutines
//  https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14


import (
    "fmt"
    "testing"
    "math"
)

const kLoopCount int = 1000000000

func MyFunction() float32 {
    return math.Pi * math.Pi
}

func Benchmark_Routine1(b *testing.B) {
    for i := 0; i < kLoopCount; i++ {
        go MyFunction()
    }
}

func Benchmark_Routine2(b *testing.B) {
    half := kLoopCount >> 1

    //done := make(chan bool)
    go func() {
        for i := 0; i < half; i++ {
            MyFunction()
        }
        //done <- true
    }()

    go func() {
        for i := 0; i < half; i++ {
            MyFunction()
        }
    }()
    fmt.Println("Bench Test Done")
}


func Benchmark_NonRoutine(b *testing.B) {
    for i := 0; i < kLoopCount; i++ {
        MyFunction()
    }
    fmt.Println("Bench Test Done")
}

func Test_NonRoutine(t *testing.T) {
    for i:=0; i<100; i++ {
        MyFunction()
    }
}

// Note: cannot run using main
func Test_simple(t *testing.T) {
    fmt.Println("Testing Goroutine")
}
