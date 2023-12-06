package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Phase string

const (
	seed        Phase = "seed"
	soil        Phase = "soil"
	fertilizer  Phase = "fertilizer"
	water       Phase = "water"
	light       Phase = "light"
	temperature Phase = "temperature"
	humidity    Phase = "humidity"
	location    Phase = "location"
)

type Seed int

type SeedRange struct {
	start  int
	length int
}

type Mapping struct {
	source      Phase
	destination Phase
	ranges      []*Range
}

func (m *Mapping) GetDest(source int) int {
	for _, r := range m.ranges {
		inRange, v := r.in(source)
		if inRange {
			return v.destination
		}
	}

	return source
}

type Range struct {
	sourceStart      int
	destinationStart int
	length           int
}

type Value struct {
	source      int
	destination int
}

func (r *Range) in(source int) (bool, *Value) {

	if source >= r.sourceStart && source < r.sourceStart+r.length {
		offset := source - r.sourceStart
		return true, &Value{
			source:      source,
			destination: r.destinationStart + offset,
		}
	}

	return false, nil
}

type Mapings map[string]*Mapping

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
	mappings := make(Mapings)
	mappingsOrder := make([]*Mapping, 0)
	//seeds := make([]Seed, 0)
	seedRanges := make([]SeedRange, 0)

	var m *Mapping
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		l = strings.Trim(l, " ")

		if strings.HasPrefix(l, "seeds: ") {
			// parse seeds
			seedsTokens, _ := strings.CutPrefix(l, "seeds: ")
			seedsNumbers := strings.Split(seedsTokens, " ")
			i := 0
			for i < len(seedsNumbers) {
				startNo, _ := strconv.Atoi(seedsNumbers[i])
				rangeLen, _ := strconv.Atoi(seedsNumbers[i+1])

				seedRanges = append(seedRanges, SeedRange{
					start:  startNo,
					length: rangeLen,
				})

				//for sNo := startNo; sNo < startNo+rangeLen; sNo++ {
				//	seeds = append(seeds, Seed(sNo))
				//}
				i += 2
			}
		} else if l == "" {
			// begining of the map
			m = &Mapping{}
		} else if !unicode.IsDigit(rune(l[0])) {
			// map name line
			mapKey, _ := strings.CutSuffix(l, " map:")

			mappingPhases := strings.Split(mapKey, "-to-")
			m.source = Phase(mappingPhases[0])
			m.destination = Phase(mappingPhases[1])
			//mappings[mapKey] = m
			mappingsOrder = append(mappingsOrder, m)
		} else if unicode.IsDigit(rune(l[0])) {
			// map range
			vals := strings.Split(l, " ")
			destStart, _ := strconv.Atoi(vals[0])
			srcStart, _ := strconv.Atoi(vals[1])
			length, _ := strconv.Atoi(vals[2])
			r := &Range{
				destinationStart: destStart,
				sourceStart:      srcStart,
				length:           length,
			}
			m.ranges = append(m.ranges, r)

		} else {
			log.Fatalf("failed on line: %s", l)
		}
	}

	//fmt.Println("Seeds: ", seeds)
	fmt.Println("Mappings order:", mappingsOrder)
	fmt.Println(mappings)

	lowest := math.MaxInt

	fmt.Println("seed ranges: ", len(seedRanges))
	wg := &sync.WaitGroup{}
	var res = make(chan int)

	for i, sr := range seedRanges {
		wg.Add(1)
		go func(no int, r SeedRange, result chan<- int, group *sync.WaitGroup) {
			defer group.Done()
			start := time.Now()
			rlow := math.MaxInt64
			fmt.Println(no, " range:", r.start, "len", r.length)
			for s := r.start; s < r.start+r.length; s++ {
				//for _, s := range seeds {
				dest := s
				for _, m := range mappingsOrder {
					//_ = mappings[mapKey]
					dest = m.GetDest(dest)
					//fmt.Printf("seed %d: phase: %s dest: %d\n", s, mapKey, dest)
				}

				if dest < rlow {
					rlow = dest
				}
			}
			fmt.Println(no, " range, took to proces: ", int(time.Since(start).Seconds()), "sec")

			result <- rlow
			fmt.Println(no, "range done!")
		}(i, sr, res, wg)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for low := range res {
		if low < lowest {
			lowest = low
		}
	}

	total = lowest

	fmt.Println(total)
	//answer part 2: 15290096
}
