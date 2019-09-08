package main

import (
	"fmt"
    "os"
    "strconv"
)

func rgb2hex(r, g, b int64) string {
    var result string

    result += fmt.Sprintf("%02X", r)
    result += fmt.Sprintf("%02X", g)
    result += fmt.Sprintf("%02X", b)

    return result
}

func rgbString(r, g, b int64) string {
    return fmt.Sprintf("(%v %v %v)", r, g, b)
}

func main() {
//arg := os.Args[3]
    //fmt.Printf("arg=%v\n", os.Args)
    if len(os.Args) != 4 {
        fmt.Println("arg: red green blue")
        fmt.Println("e.g rgb2hex 123 100 44")
        return;
    }


    r, _ := strconv.ParseInt(os.Args[1], 10, 64)
    g, _ := strconv.ParseInt(os.Args[2], 10, 64)
    b, _ := strconv.ParseInt(os.Args[3], 10, 64)
    //
    // r = int(r)
    // g = int(g)
    // b = int(b)

    //
    //fmt.Printf("%T %v %v\n", r, r, err)

    hex := rgb2hex(r, g, b)
    rgb := rgbString(r, g, b)

	fmt.Printf("rgb:%v hex:%v\n", rgb, hex)
}
