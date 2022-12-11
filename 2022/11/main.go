package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Id          int
	Items       []int
	Op          func(int) int
	Divisible   int
	True        int
	False       int
	Inspections int
}

var monkeys = make(map[int]*Monkey)

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}

	var monkey *Monkey

	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//if
		if strings.HasPrefix(line, "Monkey ") {
			monkeyTokens := strings.Split(line, ":")

			monkeyId, _ := strconv.Atoi(string(monkeyTokens[0][len(monkeyTokens[0])-1]))
			monkey = &Monkey{Id: monkeyId}

			monkeys[monkeyId] = monkey
			continue
		}

		if strings.Contains(line, "Starting items:") {
			itemsTokens := strings.Split(line, ":")

			items := strings.Split(itemsTokens[1], ", ")
			for _, item := range items {
				i, _ := strconv.Atoi(strings.TrimSpace(item))
				monkey.Items = append(monkey.Items, i)
			}
			continue
		}

		if strings.Contains(line, "Operation:") {
			ops := strings.ReplaceAll(line, "Operation: new = ", "")

			monkey.Op = parseOps(ops)
			continue
		}

		if strings.Contains(line, "Test:") {
			divisible := strings.ReplaceAll(line, "Test: divisible by ", "")
			i, _ := strconv.Atoi(strings.TrimSpace(divisible))
			monkey.Divisible = i
			continue
		}

		if strings.Contains(line, "If true: throw to monkey ") {
			divisible := strings.ReplaceAll(line, "If true: throw to monkey ", "")
			trueMonkey, _ := strconv.Atoi(strings.TrimSpace(divisible))
			monkey.True = trueMonkey
			continue
		}

		if strings.Contains(line, "If false: throw to monkey") {
			divisible := strings.ReplaceAll(line, "If false: throw to monkey ", "")
			falseMonkey, _ := strconv.Atoi(strings.TrimSpace(divisible))
			monkey.False = falseMonkey
			continue
		}

	}

	var monkeySlice = make([]*Monkey, len(monkeys))

	for k, v := range monkeys {
		fmt.Printf("%d, %v\n", k, v)
		monkeySlice[k] = v
	}

	for i := 1; i <= 20; i++ {
		fmt.Println("Round ", i)
		playRound(i, monkeySlice)
		display(monkeySlice)
		//break
	}

	sort.Slice(monkeySlice, func(i, j int) bool {
		return monkeySlice[i].Inspections > monkeySlice[j].Inspections
	})

	fmt.Printf("moneky business: %d\n", monkeySlice[0].Inspections*monkeySlice[1].Inspections)
}

func display(monkeySlice []*Monkey) {
	for _, m := range monkeySlice {
		fmt.Printf("Moneky %d: %v inspected %d\n", m.Id, m.Items, m.Inspections)
	}
}

func playRound(i int, slice []*Monkey) {
	for _, m := range slice {
		for {
			if len(m.Items) > 0 {
				item := m.Items[0]
				level := m.Op(item)
				level = level / 3
				var receivingMonkey int
				if level%m.Divisible == 0 {
					receivingMonkey = m.True
				} else {
					receivingMonkey = m.False
				}
				monkeys[receivingMonkey].Items = append(monkeys[receivingMonkey].Items, level)
				m.Inspections++
			}

			if len(m.Items) > 0 {
				m.Items = m.Items[1:]
			} else {
				break
			}
		}
	}
}

func parseOps(ops string) func(int) int {
	if strings.Contains(ops, "+") {
		tokens := strings.Split(ops, "+")
		i, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))
		return func(old int) int { return old + i }
	}
	if strings.Contains(ops, "*") {
		tokens := strings.Split(ops, "*")
		token := strings.TrimSpace(tokens[1])
		if token == "old" {
			return func(old int) int { return old * old }
		}
		i, _ := strconv.Atoi(token)
		return func(old int) int { return old * i }
	}

	panic(fmt.Sprintf("unsupported ops: %s", ops))
}
