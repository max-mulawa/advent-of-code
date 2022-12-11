package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Move struct {
	direction Direction
	distance  int
}

type Direction rune

const (
	right Direction = 'R'
	left  Direction = 'L'
	up    Direction = 'U'
	down  Direction = 'D'
)

var visitedHead = []Point{}
var visitedTail = []Point{}

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}

	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}

	count := 0
	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		move, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		move = strings.TrimSpace(move)
		if move == "" {
			break
		}
		count++

		tokens := strings.Split(move, " ")
		dir := Direction(tokens[0][0])
		distance, _ := strconv.Atoi(tokens[1])

		// if distance >= 10 {
		// 	fmt.Printf("%d", distance)
		// }

		for distance > 0 {
			// audit
			visitedHead = append(visitedHead, head)
			visitedTail = append(visitedTail, tail)

			headMove := Move{direction: dir, distance: 1}

			head, tail = makeMove(head, tail, headMove)
			distance--
		}

	}

	fmt.Printf("unique visited tail points: %d ", countDistinctPoints(visitedTail))
	fmt.Printf("executed %d moves", count)

}

func makeMove(head, tail Point, headMove Move) (newHead Point, newTail Point) {
	newTail = tail

	switch headMove.direction {
	case up:
		newHead = Point{x: head.x, y: head.y + 1}
	case down:
		newHead = Point{x: head.x, y: head.y - 1}
	case left:
		newHead = Point{x: head.x - 1, y: head.y}
	case right:
		newHead = Point{x: head.x + 1, y: head.y}
	default:
		panic(fmt.Sprintf("wrong direction: %d", headMove.direction))
	}

	newDist := distance(tail, newHead)
	if newDist > math.Sqrt2 {
		newTail = head
	}

	return newHead, newTail
}

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(float64((p2.x-p1.x)*(p2.x-p1.x)) + float64((p2.y-p1.y)*(p2.y-p1.y)))
}

func countDistinctPoints(points []Point) int {
	uniq := make(map[Point]bool)

	for _, p := range points {
		uniq[p] = true
	}

	return len(uniq)
}
