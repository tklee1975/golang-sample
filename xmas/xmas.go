package main

import (
	"fmt"
    "strings"
)

const kMargin int = 10

// 10 - 1
//
func ShowTri(triStr string) {
    margin := kMargin - (len(triStr)+1)/2
    marginStr := strings.Repeat(" ", margin)
    fmt.Printf("%v%v\n", marginStr, triStr)

    //fmt.Printf(fmtString, triStr)
}

func Triangle(size int) {
    for i :=1; i<=size; i+=2 {
        if size > 5 && i == 1 {
            continue
        }

        stars := strings.Repeat("*", i)
        ShowTri(stars)
    }
}

func Tree(height int) {
    for i :=5; i<=height; i+=2 {
        Triangle(i)
    }
}


func main() {
	//fm
    Tree(10)
}
