package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Value struct {
	v            int
	next         *Value
	prev         *Value
	diff         *Value
	extrapolated bool
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

	total := 0
	s := bufio.NewScanner(f)

	var results []*Value

	for s.Scan() {
		l := strings.Trim(s.Text(), " ")
		if l == "" {
			continue
		}

		resultsTokens := removeEmpty(strings.Split(l, " "))
		var firstRes *Value
		var prev *Value
		for _, v := range resultsTokens {
			val, _ := strconv.Atoi(v)
			valItem := &Value{
				v:    val,
				prev: prev,
			}
			if prev != nil {
				prev.next = valItem
			}
			prev = valItem

			if firstRes == nil {
				firstRes = valItem
			}
		}

		results = append(results, firstRes)
	}

	for _, r := range results {
		f := r
		t := f.extrapolate()
		total += t
		fmt.Printf("total for this results: %d\n", t)
		f.print(0)
		fmt.Printf("\n\n")
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

func (v *Value) print(level int) {
	f := v
	d := f.diff
	for f != nil {
		if f.extrapolated {
			fmt.Printf("(%d) ", f.v)
		} else {
			fmt.Printf("%d ", f.v)
		}

		f = f.next
	}
	if d != nil {
		fmt.Printf("\n")
		fmt.Printf(strings.Repeat(" ", level+1))
		d.print(level + 1)
	}
}

func (v *Value) extrapolate() int {
	exValDiff, _ := v.addDiffs()

	// traverse to the the end of the list
	f := v
	var l *Value
	for f != nil {
		l = f
		f = f.next
	}

	l.next = &Value{
		v:            l.v + exValDiff,
		prev:         l,
		extrapolated: true,
	}
	return l.next.v
}

func (v *Value) addDiffs() (int, int) {
	allZeros := true
	// traverse to the the end of the list
	f := v
	var l *Value
	for f != nil {
		l = f
		f = f.next
	}

	// calculate diff
	var nextDiff *Value
	var lastDiff *Value
	//last := l
	for l != nil {
		if l.prev == nil {
			l.diff = nextDiff
			break
		}

		diffVal := l.v - l.prev.v
		if diffVal != 0 {
			allZeros = false
		}
		l.diff = &Value{
			v: diffVal,
		}
		if nextDiff != nil {
			l.diff.next = nextDiff
			nextDiff.prev = l.diff
		} else {
			lastDiff = l.diff
		}
		nextDiff = l.diff
		l = l.prev
	}

	exVal := 0
	total := 0
	exValDiff := 0
	if !allZeros {
		exValDiff, total = nextDiff.addDiffs()
		exVal = lastDiff.v + exValDiff
	}
	lastDiff.next = &Value{
		v:            exVal,
		prev:         l,
		extrapolated: true,
	}
	return exVal, total + exVal + exValDiff
}
