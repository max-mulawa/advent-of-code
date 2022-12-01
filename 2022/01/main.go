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
	// payload, err := os.ReadFile("test-basic.txt")
	if err != nil {
		panic(err)
	}

	elfMaxCalories := 0
	elfIdx := 0
	elfMaxCalIdx := 0
	topElves := []int{}

	r := bufio.NewReader(bytes.NewReader(payload))
	elfCalories := 0
	for {
		calVal, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		calVal = strings.TrimSpace(calVal)

		if calVal == "" {
			elfIdx++
			if elfCalories > elfMaxCalories {
				elfMaxCalIdx = elfIdx
				elfMaxCalories = elfCalories
			}

			topElves = AtLeastOne(topElves, elfCalories)
			if len(topElves) > 3 {
				topElves = topElves[(len(topElves) - 3):]
			}

			elfCalories = 0
			continue
		}

		cal, _ := strconv.Atoi(calVal)
		elfCalories += cal
	}

	fmt.Printf("elf no %d has highest calories %d\n", elfMaxCalIdx, elfMaxCalories)
	fmt.Println("top 3 elves with cals:", topElves)
	totalTop3 := 0
	for i := 0; i < len(topElves); i++ {
		totalTop3 += topElves[i]
	}
	fmt.Println("top 3 elves total cals:", totalTop3)
}

func AtLeastOne(vals []int, testVal int) []int {
	if len(vals) == 0 {
		return []int{testVal}
	}
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] <= testVal {
			prev := make([]int, len(vals[:i+1]))
			copy(prev, vals[:i+1])
			x := append(prev, testVal)
			y := append(x, vals[i+1:]...)
			return y
		}
	}
	return vals
}
