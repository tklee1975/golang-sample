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
    "testing"
    "fmt"
    // "time"
    // "math/rand"
)


func TestCard(t *testing.T) {
    //card := Card {suite:Spade, rank:King};
    card := Card {suite:Spade, rank:2};

    fmt.Printf("card      : %v\n", card.string())
}


func TestCardList(t *testing.T) {
    //card := Card {suite:Spade, rank:King};
    cardList := createCardDeck()

    ShowCardDeck(cardList)
}

func TestShuffle(t *testing.T) {
    //card := Card {suite:Spade, rank:King};
    cardList := createCardDeck()

    ShuffleCard(cardList)
    ShowCardDeck(cardList)
}
