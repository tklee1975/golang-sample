package channel_test

//  use
//  go test channel_test.go
//  go benchmark goroutine_test.go
// to test
//
// Reference:
//  https://gobyexample.com/goroutines
//  https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14


import (
    "fmt"
    // "bufio"
    // "os"
    "testing"
)

func Test_SimpleMessage(t *testing.T) {
    message := make(chan string)

    go func() {
        message <- "ping"
    }()

    fmt.Printf("Wait for message\n")
    msg := <-message

    fmt.Printf("%v\n", msg)
}


//
// func Test_Input(t *testing.T) {
//     //message := make(chan string)
//
//     go func() {
//         scanner := bufio.NewScanner(os.Stdin)
//         for scanner.Scan() {
//             fmt.Println(scanner.Text())
//         }
//
//     }()
//
//     // fmt.Printf("Wait for message\n")
//     // msg := <-message
//     //
//     // fmt.Printf("%v\n", msg)
// }
