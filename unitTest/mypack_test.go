package mypack

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

func TestSimple(t *testing.T) {
    fmt.Println("Testing")
}

func Add(a int, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    if Add(1, 2) == 4 {
        t.Log("Add.Add Pass")
    } else {
        t.Errorf("Add.Add Fail")
    }
}

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}


func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    t.Log("Time Consume")
}

func TestErrorf(t *testing.T) {
    var a = 1
    var b = 2

    if a + b == 3 {
        t.Log("Success!")
    } else {
        t.Errorf("Testing error!!!")
    }
}
