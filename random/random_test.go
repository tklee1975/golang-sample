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
    "math/rand"
    "time"
)

func TestRandom1(t *testing.T) {
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 10; i++ {
        value := rand.Intn(100)

        fmt.Printf("%v ", value)
    }
}

func TestRandom2(t *testing.T) {
    fmt.Printf("time=%v\n", time.Now().UTC().UnixNano())

    rand.Seed(time.Now().UTC().UnixNano())

    for i := 0; i < 10; i++ {
        value := rand.Intn(100)

        fmt.Printf("%v ", value)
    }
}
