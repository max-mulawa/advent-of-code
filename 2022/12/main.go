package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	//prev = make(map[Vertex]Vertex)
	//dist = make(map[Vertex]int)
	//unvisited = make(map[Vertex]bool)

	input = [][]rune{}
	edges = make(map[Vertex][]Edge)
)

type Edge struct {
	v      Vertex
	weight int
}

type Vertex struct {
	x, y int
}

var (
	startV Vertex
	destV  Vertex

	inifinity = 1000000
)

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}

	rowCount := 0
	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		row, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		row = strings.TrimSpace(row)
		if row == "" {
			break
		}

		vertecies := make([]rune, len(row))
		for col, v := range row {
			if v == 'S' {
				v = 'a'
				startV = Vertex{x: col, y: rowCount}
				// destV = Vertex{x: col, y: rowCount}
			} else if v == 'E' {
				v = 'z'
				destV = Vertex{x: col, y: rowCount}
				// startV = Vertex{x: col, y: rowCount}
			}

			vertecies[col] = v
		}

		input = append(input, vertecies)

		rowCount++
	}

	maxX := len(input[0]) - 1
	maxY := len(input) - 1

	var vertecies = []Vertex{}

	render(func(vertex Vertex, r rune) rune {
		vertecies = append(vertecies, vertex)
		createEdges(vertex, maxX, maxY)
		return r
	})

	// checkV := destV
	// for _, e := range edges[checkV] {
	// 	fmt.Printf("%v to %v edge\n", checkV, e.v)
	// }

	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode

	//minPath := calcMinPath(startV, vertecies)

	s := sync.WaitGroup{}
	res := sync.Map{}

	render(func(vertex Vertex, r rune) rune {

		if r == 'a' {
			s.Add(1)
			go func() {
				defer s.Done()
				path := calcMinPath(vertex, vertecies)
				res.Store(vertex, path)
			}()
		}
		return r
	})
	s.Wait()

	minPath := inifinity
	res.Range(func(key, value any) bool {
		if value.(int) < minPath {
			minPath = value.(int)
		}
		return true
	})

	// render(func(vertex Vertex, r rune) rune {
	// 	if dist[vertex] != inifinity {
	// 		return '0'
	// 	}
	// 	return r
	// })

	fmt.Printf("minimal path: %d", minPath)
}

func calcMinPath(start Vertex, vertecies []Vertex) int {
	unvisited := make(map[Vertex]bool)
	dist := make(map[Vertex]int)
	for _, v := range vertecies {
		unvisited[v] = true
	}

	for k := range unvisited {
		dist[k] = inifinity
	}

	dist[start] = 0
	for len(unvisited) > 0 {
		u := getMin(dist, unvisited)
		delete(unvisited, u)

		//fmt.Printf("current min node: %v with dist: %d for height: %c \n", u, dist[u], input[u.y][u.x])

		if u == destV {
			break
		}

		for _, edge := range edges[u] {
			if !unvisited[edge.v] {
				continue
			}

			alt := dist[u] + edge.weight
			if alt < dist[edge.v] {
				dist[edge.v] = alt
				//prev[edge.v] = u
			}
		}
	}

	// render(func(vertex Vertex, r rune) rune {
	// 	if dist[vertex] != inifinity {
	// 		return '0'
	// 	}
	// 	return r
	// })

	return dist[destV]
}

func render(perV func(Vertex, rune) rune) {
	for i, r := range input {
		for j, h := range r {
			vertex := Vertex{
				x: j,
				y: i,
			}
			c := perV(vertex, h)
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func getMin(pathDist map[Vertex]int, unvisited map[Vertex]bool) Vertex {
	var minDistVert *Vertex
	var minDist int = inifinity
	for k := range unvisited {
		v := k
		distance := pathDist[k]

		if minDistVert != nil {
			if minDist > distance {
				minDistVert = &v
				minDist = distance
			}
		} else {
			minDistVert = &v
			minDist = distance
		}
	}
	if minDistVert == nil {
		panic("getMin failed")
	}

	return *minDistVert
}

func createEdges(vertex Vertex, maxX, maxY int) {
	for _, xdiff := range []int{-1, 1} {
		x := vertex.x + xdiff
		y := vertex.y
		createEdge(vertex, x, y, maxX, maxY)
	}

	for _, ydiff := range []int{-1, 1} {
		x := vertex.x
		y := vertex.y + ydiff
		createEdge(vertex, x, y, maxX, maxY)
	}
}

func createEdge(vertex Vertex, x, y, maxX, maxY int) {
	height := input[vertex.y][vertex.x]
	if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
		v := input[y][x]

		diff := int(height) - int(v)
		if diff >= 0 || diff == -1 {
			otherV := Vertex{x: x, y: y}
			edge := Edge{
				v:      otherV,
				weight: 1,
			}
			edges[vertex] = append(edges[vertex], edge)
		}
	}
}
