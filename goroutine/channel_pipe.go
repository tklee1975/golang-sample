package main

//  use
//  go test channel_test.go
//  go benchmark goroutine_test.go
// to test
//
// Reference:
//  https://gobyexample.com/goroutines
//  https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14
// https://gobyexample.com/channel-directions


import (
    "fmt"
//    "strings"
    //"testing"
)



func PipeInOut(in <-chan string, out chan<- string) {
    msg := <-in
    msg = fmt.Sprintf("%v len=%d", msg, len(msg))
    out <- msg

}

func main() {
    in := make(chan string, 1)
    out := make(chan string, 1)

    in <- "Testing"
    PipeInOut(in, out)
    fmt.Printf("out=%v\n", <-out)
}
