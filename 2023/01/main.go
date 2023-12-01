package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//f, err := os.OpenFile("base2.txt", os.O_RDONLY, 0x664)
	f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		l := s.Text()
		arr := []rune(l)
		len := len(l)
		first := 0
		last := 0
		for i := 0; i < len; i++ {
			r := arr[i]
			if unicode.IsDigit(r) && first == 0 {
				first, _ = strconv.Atoi(string([]rune{r}))
				first = first * 10
			} else if first == 0 {
				d, err := parseDigit(string(arr[i:len]), true)
				if err == nil {
					first = d * 10
				}
			}

			r = arr[len-1-i]
			if unicode.IsDigit(r) && last == 0 {
				last, _ = strconv.Atoi(string([]rune{r}))
			} else if last == 0 {
				d, err := parseDigit(string(arr[0:len-i]), false)
				if err == nil {
					last = d
				}
			}

		}

		//fmt.Println(l, first, last)
		total += first + last
	}

	fmt.Println(total)
}

func parseDigit(val string, prefix bool) (int, error) {

	for txt, d := range digits {
		if prefix && strings.HasPrefix(val, txt) {
			return d, nil
		} else if !prefix && strings.HasSuffix(val, txt) {
			return d, nil
		}
	}
	return -1, fmt.Errorf("digit not found")
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Part1() {
	//f, err := os.OpenFile("base.txt", os.O_RDONLY, 0x664)
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		l := s.Text()
		arr := []rune(l)
		len := len(l)
		first := 0
		last := 0
		for i := 0; i < len; i++ {
			r := arr[i]
			if unicode.IsDigit(r) && first == 0 {
				first, _ = strconv.Atoi(string([]rune{r}))
				first = first * 10
			}

			r = arr[len-1-i]
			if unicode.IsDigit(r) && last == 0 {
				last, _ = strconv.Atoi(string([]rune{r}))
			}
		}

		//fmt.Println(l, first, last)
		total += first + last
	}

	fmt.Println(total)
}
