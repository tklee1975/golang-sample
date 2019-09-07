package main

import (
    "fmt"
    "math/rand"
    "time"
)


func getArray(size int) []int {
    var array []int

    for i:=0; i<size; i++ {
        array = append(array, (i+1))
    }

    return array;
}

func swap(a, b int) (int, int) {
    return b, a
}

func shuffleList(data []int) {
    var dataSize = len(data)


    rand.Seed(time.Now().UnixNano())

    //r := rand.New(rand.NewSource(50))

    for i :=0; i<30; i++ {
        var pos1 = rand.Intn(dataSize)
        var pos2 = rand.Intn(dataSize)

        data[pos1], data[pos2] = swap(data[pos1], data[pos2])
    }
    //rand.Seed(time.Now().UnixNano())
    // a = data[0]
    // data[0] = data[1]
    // data[1] = a
}

func shuffleListDemo() {
    var array []int

    array = getArray(10)
    shuffleList(array)

    fmt.Printf("%v\n", array);
}

func main() {
    shuffleListDemo()
}
