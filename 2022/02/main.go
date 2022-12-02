package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

func main() {
	payload, err := os.ReadFile("input.txt")
	//payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}
	totalRes := 0

	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		play, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		play = strings.TrimSpace(play)
		if play == "" {
			break
		}

		actions := strings.Split(play, " ")
		oponentAction := actions[0]
		//myAction := decode(actions[1])
		myAction := decodeV2(actions[1], oponentAction)
		_, myRes := evaluate(oponentAction, myAction)
		totalRes += myRes
	}

	fmt.Printf("my res: %d", totalRes)
}

func evaluate(action1 string, action2 string) (int, int) {
	base1 := bonus(action1)
	base2 := bonus(action2)
	res1, res2 := compare(action1, action2)
	return res1 + base1, res2 + base2
}

var (
	beats = map[string]map[string]bool{
		Rock: {
			Scissors: true,
		},
		Scissors: {
			Paper: true,
		},
		Paper: {
			Rock: true,
		},
	}

	allActions = []string{
		Rock,
		Scissors,
		Paper,
	}
)

func compare(action1 string, action2 string) (int, int) {
	if action1 == action2 {
		return 3, 3
	}

	beatsAction := beats[action1][action2]
	if beatsAction {
		return 6, 0
	}
	return 0, 6
}

func bonus(action string) int {
	switch action {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic(fmt.Sprintf("wrong action: %q", action))
	}
}

func decodeV2(action string, compAction string) string {
	switch action {
	case "X":
		return firstKey(beats[compAction])
	case "Y":
		return compAction
	case "Z":
		return winAction([2]string{compAction, firstKey(beats[compAction])})
	default:
		panic(fmt.Sprintf("wrong action: %q", action))
	}
}

func winAction(otherActions [2]string) string {
	other := make(map[string]bool)

	for _, o := range otherActions {
		other[o] = true
	}

	for _, a := range allActions {
		if !other[a] {
			return a
		}
	}
	panic("couldn't get winning action")
}

func firstKey(elems map[string]bool) string {
	for k := range elems {
		return k
	}
	panic("nothing to return")
}

func decode(action string) string {
	switch action {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic(fmt.Sprintf("wrong action: %q", action))
	}
}
