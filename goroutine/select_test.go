package select_test

//  use
//  go test channel_test.go
//  go benchmark goroutine_test.go
// to test
//
// Reference:
//  https://gobyexample.com/select
//  https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14


import (
    "fmt"
    // "bufio"
    // "os"
    "testing"
    "time"
    "runtime"
)

func SimpleSleeper(sleepTime int) {
    time.Sleep(time.Second * time.Duration(sleepTime))
}

func ChannelSleeper(sleepTime int, name string, chFlag chan string) {
    fmt.Printf("%v: Start sleeping\n", name)
    SimpleSleeper(sleepTime)
    chFlag <- fmt.Sprintf("%v done", name)
}

func Test_Simple(t *testing.T) {
    // message := make(chan string)
    //
    // go func() {
    //     message <- "ping"
    // }()

    fmt.Printf("Start Sleeping\n")

    SimpleSleeper(1)

    fmt.Printf("Program Done\n")
}


func Test_Select(t *testing.T) {
    // message := make(chan string)
    //
    // go func() {
    //     message <- "ping"
    // }()

    ch1 := make(chan string)
    ch2 := make(chan string)

    fmt.Printf("Start Select\n")

    fmt.Printf("Go routine Count=%d\n", runtime.NumGoroutine())

    go ChannelSleeper(5, "Chan1", ch1)
    go ChannelSleeper(10, "Chan2", ch2)

    fmt.Printf("Go routine Count=%d\n", runtime.NumGoroutine())

    // Start waiting the channels
    for i:=0; i<2; i++ {        // kencoder: will show 'deadlock' when the number of count > 2
        fmt.Printf("Loop %v Start\n", i)
        select {
        case msg1 := <-ch1:
            fmt.Printf("msg1: %v\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("msg2: %v\n", msg2)
        }
        fmt.Printf("Loop %v End routine=%d\n", i, runtime.NumGoroutine())
    }

    fmt.Printf("Program Done\n")
}
