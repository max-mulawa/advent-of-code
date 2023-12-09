package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cStrenght = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	//'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type HandScore int

const (
	High HandScore = iota
	Pair
	TwoPairs
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

type Cards struct {
	cards string
}

func (c *Cards) score() HandScore {
	res := make(map[int]int)
	countJ := 0
	maxCard := -1
	maxCount := -1
	for _, card := range c.cards {
		if card == 'J' {
			countJ++
			continue
		}
		s := cStrenght[card]
		if _, ok := res[s]; !ok {
			res[s] = 1
		} else {
			res[s] += 1
		}
		v := res[s]
		if v > maxCount {
			maxCard = s
			maxCount = v
		}

	}

	if countJ > 0 {
		if len(res) == 0 {
			maxCard = cStrenght['J']
			res[maxCard] = 0
		}
		res[maxCard] += countJ
	}

	uniq := len(res)
	if uniq == 1 {
		return FiveKind
	}

	if uniq == 5 {
		return High
	}

	if uniq == 4 {
		return Pair
	}

	if uniq == 3 {
		// TwoPairs
		// Three
		for _, v := range res {
			if v == 2 {
				return TwoPairs
			} else if v == 3 {
				return ThreeKind
			}
		}
	}

	if len(res) == 2 {
		// FourKind or FullHouse
		for _, v := range res {
			if v == 3 || v == 2 {
				return FullHouse
			}
			return FourKind

		}
	}
	return -1
}

// compare
// evaluate

type Hand struct {
	c     *Cards
	score HandScore
	bid   int
	rank  int
}

func NewHand(cards string, bid int) *Hand {
	c := &Cards{
		cards: cards,
	}
	h := &Hand{
		c:     c,
		score: c.score(),
		bid:   bid,
	}

	return h
}

func (h *Hand) less(h2 *Hand) bool {
	if h.score == h2.score {
		h1Runes := []rune(h.c.cards)
		h2Runes := []rune(h2.c.cards)
		for i, c := range h1Runes {
			if cStrenght[c] == cStrenght[h2Runes[i]] {
				continue
			} else {
				return cStrenght[c] < cStrenght[h2Runes[i]]
			}
		}
		return false
	}

	return h.score < h2.score
}

// score
// compare

func main() {
	//f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0x664)
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base2.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	total := 0
	var hands []*Hand

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.Trim(s.Text(), " ")
		handTokens := removeEmpty(strings.Split(l, " "))
		handCards := strings.Trim(handTokens[0], " ")
		bid, _ := strconv.Atoi(handTokens[1])

		hands = append(hands, NewHand(handCards, bid))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].less(hands[j])
	})

	for i, h := range hands {
		h.rank = i + 1
		fmt.Printf("hand: %s with %d score %d bid, rank: %d\n", h.c.cards, h.score, h.bid, h.rank)
		total += h.rank * h.bid
	}

	fmt.Println(total)

}

func removeEmpty(a []string) []string {
	var b []string

	for _, s := range a {
		if strings.Trim(s, "") != "" {
			b = append(b, strings.Trim(s, ""))
		}
	}

	return b
}
