package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	f, basePos := floor(input)

	fmt.Println("result floor:", f, "first basement entry: ", basePos)
}

func floor(input []byte) (floor int, basementPos int) {
	for idx, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor < 0 && basementPos == 0 {
			basementPos = idx + 1
		}
	}
	return floor, basementPos
}
