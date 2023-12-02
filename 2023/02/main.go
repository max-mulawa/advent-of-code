package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	b, r, g int
	no      int
}

func NewGame(l string) *game {
	g := game{}

	tokens := strings.Split(l, ":")
	// game number
	gNoToken := tokens[0]
	g.no, _ = strconv.Atoi(strings.Replace(gNoToken, "Game ", "", 1))

	// games
	gPlaysTokens := tokens[1]

	plays := strings.Split(gPlaysTokens, ";")
	for _, p := range plays {

		cubesTokens := strings.Split(p, ",")

		for _, c := range cubesTokens {
			cubes := strings.Split(strings.Trim(c, " "), " ")
			count, _ := strconv.Atoi(cubes[0])
			switch strings.Trim(cubes[1], " ") {
			case "red":
				if count > g.r {
					g.r = count
				}
			case "green":
				if count > g.g {
					g.g = count
				}
			case "blue":
				if count > g.b {
					g.b = count
				}
			default:
				log.Fatalf("cubes color parsing error: %s", cubes[1])
			}

		}
	}

	return &g
}

func (g *game) valid() bool {
	if g.r <= 12 && g.g <= 13 && g.b <= 14 {
		return true
	}
	return false
}

func (g *game) val() int {
	return g.r * g.g * g.b
}

func main() {
	f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base2.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		l := s.Text()
		g := NewGame(l)
		//if g.valid() {
		//total += g.no
		total += g.val()
		//}
	}

	fmt.Println(total)

}
