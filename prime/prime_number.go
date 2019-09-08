package main

/**
 *
 */
import (
	"fmt"
)

func PrimeList(rangeSize int) [] int {
    var result [] int

    var flag = make([]int, rangeSize+1)
    //var flag [rangeSize] int

    // Marking the number which is a product of something
    for i:=2; i<=rangeSize; i++ {
        max := (rangeSize+1) / i

        for j:=2; j<=max; j++ {
            product := i * j

            //fmt.Printf("%v i=%v j=%v\n", product, i, j)

            if product <= rangeSize {
                flag[product] = 1
            }
        }
    }

    for i:=1; i<=rangeSize; i++ {
        if flag[i] == 0 {
            result = append(result, i)
        }
    }
    //fmt.Printf("%v\n", flag)

    return result;
}

func main() {

    array := PrimeList(500)

    fmt.Printf("%v\n", array)
}
