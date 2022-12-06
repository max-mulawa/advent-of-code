package main

import (
	"fmt"
)

func main() {
	//dataStream := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	//dataStream := "nppdvjthqldpwncqszvftbrmjlhg"
	//dataStream := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	// dataStream := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	dataStream := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	//distinctCnt := 4
	distinctCnt := 14
	for i := 0; i < len(dataStream)-distinctCnt; i++ {
		if isUnique(dataStream[i : i+distinctCnt]) {
			fmt.Printf("first marker after character: %d\n", i+distinctCnt)
			break
		}

	}
}

func isUnique(s string) bool {
	exists := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if !exists[s[i]] {
			exists[s[i]] = true
		} else {
			return false
		}
	}
	return true
}
