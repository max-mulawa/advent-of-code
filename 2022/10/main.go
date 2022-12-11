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

	registry := 1
	cycle := 0
	cycles := []int{20, 60, 100, 140, 180, 220}
	total := 0

	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		instruction, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		instruction = strings.TrimSpace(instruction)
		if instruction == "" {
			break
		}

		if instruction == "noop" {
			cycle++
			if getCycle(cycle, cycles) > 0 {
				total += cycle * registry
				//fmt.Printf("at %d cycle registry val was %d\n", cycle, registry)
			}
			print(cycle, registry)
			continue
		}

		tokens := strings.Split(instruction, " ")

		value, _ := strconv.Atoi(tokens[1])
		for i := 0; i < 2; i++ {
			cycle++
			if getCycle(cycle, cycles) > 0 {
				total += cycle * registry
				//fmt.Printf("at %d cycle registry val was %d\n", cycle, registry)
			}
			print(cycle, registry)
		}
		registry += value
	}

	fmt.Printf("total is %d", total)
}

func getCycle(current int, c []int) int {
	for _, cycle := range c {
		if current == cycle {
			return cycle
		}
	}
	return 0
}

func print(cycle int, regVal int) {
	crtPos := (cycle - 1) % 40
	if crtPos >= regVal-1 && crtPos <= regVal+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	if cycle%40 == 0 {
		fmt.Printf("\n")
	}
}
