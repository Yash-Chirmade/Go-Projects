package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

//Function to create deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"ace", "two", "three", "four", "five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, (value + " of " + suit))
		}
	}
	return cards

} //Function to convert the slice of card to a single string
func (d deck) tostring() string {
	return strings.Join(d, ",")
}

//Function to create two hands
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

//Function to print card in deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Function to store data to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.tostring()), 0666)

}

//Function to read the file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
	//converting the byte file back to slice of strings
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//Function to shuffle the cards
func (d deck) shuffle() {
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
