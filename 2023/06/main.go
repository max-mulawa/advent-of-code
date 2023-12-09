package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	distanceToBeat int
	time           int
}

func (r *Race) WinningOpts() int {
	// minV  mil/ms
	// distance = V * s
	options := 0
	v := 1
	for t := r.time - 1; t > 0; t-- {
		if r.distanceToBeat < t*v {
			options++
		}
		v += 1
	}
	return options
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

	total := 1
	var races []*Race
	var timeTokens []string
	var distanceTokens []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := s.Text()
		l = strings.Trim(l, " ")

		if strings.HasPrefix(l, "Time: ") {
			// parse time
			timeLine, _ := strings.CutPrefix(l, "Time: ")
			timeLine = strings.Replace(timeLine, " ", "", -1)
			timeTokens = removeEmpty(strings.Split(timeLine, " "))
		} else if strings.HasPrefix(l, "Distance: ") {
			// parse distance
			distanceLine, _ := strings.CutPrefix(l, "Distance: ")
			distanceLine = strings.Replace(distanceLine, " ", "", -1)
			distanceTokens = removeEmpty(strings.Split(distanceLine, " "))
		} else {
			log.Fatalf("Invalid file format: %s", l)
		}
	}

	if len(distanceTokens) != len(timeTokens) {
		log.Fatalf("distance count %d doesn't match time count %d", len(distanceTokens), len(timeTokens))
	}

	tokenCount := len(timeTokens)

	for i := 0; i < tokenCount; i++ {
		time, _ := strconv.Atoi(timeTokens[i])
		distance, _ := strconv.Atoi(distanceTokens[i])

		races = append(races, &Race{
			time:           time,
			distanceToBeat: distance,
		})
	}

	for _, r := range races {
		opt := r.WinningOpts()
		fmt.Println(r, "has", opt)
		total *= opt
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
