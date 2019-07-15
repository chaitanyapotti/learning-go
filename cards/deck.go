package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of string
type deck []string

// (d deck) is called a receiver of a function
// Any variable of type deck gets access to print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	suites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suite := range suites {
		for _, number := range numbers {
			cards = append(cards, number+" Of "+suite)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	// var finalString string
	// for _, deckItem := range d {
	// 	finalString += deckItem + ","
	// }
	// return finalString

	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a new deck
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomNum := r.Intn(len(d) - 1)
		// temp := d[randomNum]
		// d[randomNum] = card
		// d[i] = temp
		d[i], d[randomNum] = d[randomNum], d[i]
	}
}
