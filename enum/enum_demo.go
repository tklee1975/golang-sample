package main

// https://yourbasic.org/golang/iota/

import (
	"fmt"
)

type Direction int
const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func directionEnum() {
    a := East       // or var a Direction = East
    b := North      // or var b = North

    fmt.Printf("a=%v b=%v\n", a, b)
}


func main() {
	//fmt.Println("Hello Go!")
	directionEnum()
}
