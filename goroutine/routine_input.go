package main

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
    "bufio"
    "os"
    "strings"
    //"testing"
)

func readInput(message chan string, done chan bool) {
    reader := bufio.NewReader(os.Stdin)

    for true {
        fmt.Print("Enter text: ")
        text, _ := reader.ReadString('\n')

        text = strings.TrimSpace(text)
        fmt.Println();
        //fmt.Printf("%v\n", text)

        if text == "exit" {
            break
        }

        message <- text
    }

    done <- true
}

func echoInput(message chan string) {
    for true {
        msg := <-message
        fmt.Printf("echo: %v\n", msg)
    }
}

func main() {
    message := make(chan string)
    done := make(chan bool)
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter text: ")
    // text, _ := reader.ReadString('\n')
    // fmt.Println(text)
    go readInput(message, done)
    go echoInput(message)

    // message := make(chan string)
    //
    // go func() {
    //     message <- "ping"
    // }()
    //
    // fmt.Printf("Wait for message\n")
    // for true {
    //     msg := <-message
    //     fmt.Printf("echo: %v\n", msg)
    // }
    <- done
    fmt.Println("Complete")
}
