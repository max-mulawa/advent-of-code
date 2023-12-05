package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	name       string
	numbers    map[int]bool
	winNumbers map[int]bool
	copies     []*card
}

func NewCard(l string) *card {
	c := card{
		numbers:    make(map[int]bool),
		winNumbers: make(map[int]bool),
	}

	tokens := strings.Split(l, ":")

	c.name = tokens[0]
	numbers := tokens[1]
	numbersTokens := strings.Split(numbers, "|")

	winingNumersTokens := strings.Split(strings.Trim(numbersTokens[0], " "), " ")
	numTokens := strings.Split(strings.Trim(numbersTokens[1], " "), " ")

	for _, n := range winingNumersTokens {
		v := strings.Trim(n, " ")
		if v == "" {
			continue
		}
		val, _ := strconv.Atoi(v)
		c.winNumbers[val] = true
	}

	for _, n := range numTokens {
		v := strings.Trim(n, " ")
		if v == "" {
			continue
		}
		val, _ := strconv.Atoi(v)
		c.numbers[val] = true
	}

	return &c
}

func (c *card) value() int {
	val := 0
	for k, _ := range c.winNumbers {
		if c.numbers[k] {
			if val == 0 {
				val = 1
			} else {
				val *= 2
			}
		}
	}
	return val
}

func (c *card) wining() string {
	val := ""
	for k, _ := range c.winNumbers {
		if c.numbers[k] {
			val += fmt.Sprintf("%d,", k)
		}
	}
	return val
}

func (c *card) winnersCount() int {
	val := 0
	for k, _ := range c.winNumbers {
		if c.numbers[k] {
			val += 1
		}
	}
	return val
}

func main() {
	//f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0x664)
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base2.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var cards []*card
	for s.Scan() {
		l := s.Text()
		cards = append(cards, NewCard(l))
	}

	total := 0

	for i, c := range cards {

		addCopies(c, i, cards)
		for _, copy := range c.copies {
			addCopies(copy, i, cards)
		}

		count := len(c.copies)
		total += count

		fmt.Println(c.name, "winners:", c.winnersCount(), " actual:", count)

		// part 1
		//v := c.value()
		//total += v
	}

	total += len(cards)

	fmt.Println(total)
}

func addCopies(c *card, i int, cards []*card) {
	count := c.winnersCount()
	for ci := i + 1; ci < min(i+1+count, len(cards)); ci++ {
		card := cards[ci]
		card.copies = append(card.copies, card)
	}
}
