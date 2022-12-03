package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}
	totalCompartmentOverlap := 0
	totalGroupsPriorities := 0
	group := []string{}
	groupCounter := 0

	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		rucksackItems, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		rucksackItems = strings.TrimSpace(rucksackItems)
		if rucksackItems == "" {
			break
		}

		group = append(group, rucksackItems)
		groupCounter++

		if groupCounter == 3 {
			totalGroupsPriorities += getGroupPriority(group)
			groupCounter = 0
			group = []string{}
		}

		totalCompartmentOverlap += getCompartmentOverlapPriority(rucksackItems)
	}

	fmt.Printf("total priority: %d\n", totalCompartmentOverlap)
	fmt.Printf("total group priority: %d\n", totalGroupsPriorities)
}

func getGroupPriority(group []string) int {
	overlappingItem := getOverlappingItem(group)
	return int(getPriority(overlappingItem))
}

func getCompartmentOverlapPriority(rucksackItems string) int {
	itemsCount := len(rucksackItems)
	firstCompartment := rucksackItems[:(itemsCount / 2)]
	secondCompartment := rucksackItems[(itemsCount / 2):]

	overlappingItems := getOverlapping(firstCompartment, secondCompartment)

	totalPriority := getItemsPriority(overlappingItems)
	return totalPriority
}

func getItemsPriority(items []byte) int {
	totalPriority := 0
	for _, item := range items {
		totalPriority += int(getPriority(item))
	}
	return totalPriority
}

func getPriority(item byte) uint8 {
	if item >= 65 && item <= 90 {
		// A-Z => 65-90 ASCII => 27->52 (shift by 38)
		return item - 38
	} else if item >= 97 && item <= 122 {
		// a-z => 97-122 ASCII => 1-26 (shit by 96)
		return item - 96
	} else {
		panic(fmt.Sprintf("cannot set priority for %d", item))
	}
}

func getOverlappingItem(rucks []string) byte {
	ruckMaps := []map[rune]bool{}
	for _, ruck := range rucks {
		ruckMaps = append(ruckMaps, toMap(ruck))
	}

	for _, i := range rucks[0] {
		inAll := true
		for _, itemMap := range ruckMaps {
			if !itemMap[i] {
				inAll = false
				break
			}
		}
		if inAll {
			return byte(i)
		}
	}

	panic("no common badge")
}

func toMap(item string) map[rune]bool {
	chars := make(map[rune]bool)
	for _, ic := range item {
		chars[ic] = true
	}
	return chars
}

func getOverlapping(firstCompartment, secondCompartment string) []byte {
	first := make(map[rune]bool)
	for _, c := range firstCompartment {
		first[c] = true
	}
	overlapping := make(map[rune]bool)
	res := []byte{}
	for _, c := range secondCompartment {
		if first[c] && !overlapping[c] {
			overlapping[c] = true
			res = append(res, byte(c))
		}
	}
	return res
}
