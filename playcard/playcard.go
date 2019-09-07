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
 *
 * https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
 */

import (
    "fmt"
    "time"
    "math/rand"
)

// Suite
type Suite int
const (
    Spade Suite = iota
    Heart
    Club
    Diamond
)

func (d Suite) string() string {
    return [...]string{"Spade", "Heart", "Club", "Diamond"}[d]
}

func (d Suite) symbol() string {
    return [...]string{"\u2660", "\u2665", "\u2663", "\u2666"}[d]
}


// Rank
type Rank int
const (
    Jack Rank = 11
    Queen Rank = 12
    King Rank = 13
)

func (d Rank) string() string {
    if Jack == d  {
        return "J";
    } else if Queen == d {
        return "Q";
    } else if King == d {
        return "K";
    } else {
        return fmt.Sprintf("%v", d);
    }
}

// Struct definition
type Card struct {
    suite Suite
    rank Rank
}



func (card *Card) string() string {
    return fmt.Sprintf("%v%-2v", card.suite.symbol(), card.rank.string());
}

func createCardDeck() []Card {
    var cardList []Card

    for i := 0; i<4; i++ {
        for j := 1; j<=13; j++ {
            suite := Suite(i)
            rank := Rank(j)
            card := Card {suite: suite, rank: rank};

            cardList = append(cardList, card)
        }
    }

    return cardList
}

func ShuffleCard(cardList []Card) {
    dataSize := len(cardList)

    rand.Seed(time.Now().UnixNano())

    //r := rand.New(rand.NewSource(50))

    for i :=0; i<30; i++ {
        var pos1 = rand.Intn(dataSize)
        var pos2 = rand.Intn(dataSize)

        tempCard := cardList[pos2]
        cardList[pos2] = cardList[pos1]
        cardList[pos1] = tempCard
    }
}

func ShowCardDeck(cardList []Card) {
    col := 0
    for _, v := range cardList {
		fmt.Printf("%s ", v.string())
        col++
        if col == 13 {
            fmt.Printf("\n")
            col = 0
        }
	}
}
