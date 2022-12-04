package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}
	totalContainedSec := 0

	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		pairsOfSections, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		pairsOfSections = strings.TrimSpace(pairsOfSections)
		if pairsOfSections == "" {
			break
		}

		sectionRanges := strings.Split(pairsOfSections, ",")
		s1 := NewSection(sectionRanges[0])
		s2 := NewSection(sectionRanges[1])
		fullyContainSection := overlap(s1, s2)
		if fullyContainSection != nil {
			totalContainedSec++
		}
	}

	fmt.Printf("total : %d\n", totalContainedSec)
}

type section struct {
	from int
	to   int
}

func NewSection(sectionAsText string) section {
	boudaries := strings.Split(sectionAsText, "-")
	from, _ := strconv.Atoi(boudaries[0])
	to, _ := strconv.Atoi(boudaries[1])
	return section{from: from, to: to}
}

func (s section) overlap(s2 section) bool {
	return (s.from <= s2.from && s.to >= s2.from) || (s.from <= s2.to && s.to >= s2.to)
}

// func (s section) fullyOverlap(s2 section) bool {
// 	return s.from <= s2.from && s.to >= s2.to && s.to >= s2.from
// }

func overlap(s1 section, s2 section) *section {
	if s1.overlap(s2) {
		return &s1
	} else if s2.overlap(s1) {
		return &s2
	} else {
		return nil
	}
}
