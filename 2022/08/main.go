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

	treeMatrix := [][]int{}
	rowCount := 0
	r := bufio.NewReader(bytes.NewReader(payload))
	for {
		treesRow, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		treesRow = strings.TrimSpace(treesRow)
		if treesRow == "" {
			break
		}

		treeRow := []int{}

		for _, t := range treesRow {
			height, _ := strconv.Atoi(string(t))
			treeRow = append(treeRow, height)
		}

		treeMatrix = append(treeMatrix, treeRow)
		rowCount++
	}

	fmt.Println(treeMatrix)

	visibleTrees := 0

	for i, row := range treeMatrix {
		for j, height := range row {
			if isTreeVisible(i, j, height, treeMatrix) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("visible trees: %d\n", visibleTrees)

	maxScore := 0
	for i, row := range treeMatrix {
		for j, height := range row {
			score := scenicViewScore(i, j, height, treeMatrix)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Printf("max visible scope: %d\n", maxScore)
}

func scenicViewScore(i, j, height int, treeMatrix [][]int) int {
	if i == 0 || i == (len(treeMatrix)-1) {
		return 0
	}
	if j == 0 || j == len(treeMatrix[i])-1 {
		return 0
	}

	treesColumn := []int{}
	for rowIdx := 0; rowIdx < len(treeMatrix); rowIdx++ {
		treesColumn = append(treesColumn, treeMatrix[rowIdx][j])
	}

	top := howManyTreesVisible(reverseArray(treesColumn[:i]), height)
	bottom := howManyTreesVisible(treesColumn[i+1:], height)

	left := howManyTreesVisible(reverseArray(treeMatrix[i][:j]), height)
	right := howManyTreesVisible(treeMatrix[i][j+1:], height)

	return left * right * top * bottom
}

func howManyTreesVisible(trees []int, height int) int {
	count := 0
	for i := 0; i < len(trees); i++ {
		count++
		if trees[i] >= height {
			return count
		}
	}
	return count
}

func isTreeVisible(i, j, height int, treeMatrix [][]int) bool {
	if i == 0 || i == (len(treeMatrix)-1) {
		return true
	}
	if j == 0 || j == len(treeMatrix[i])-1 {
		return true
	}

	treesColumn := []int{}
	for rowIdx := 0; rowIdx < len(treeMatrix); rowIdx++ {
		treesColumn = append(treesColumn, treeMatrix[rowIdx][j])
	}

	topBottom := isVisible(treesColumn[:i], height) || isVisible(treesColumn[i+1:], height)
	leftRigh := isVisible(treeMatrix[i][:j], height) || isVisible(treeMatrix[i][j+1:], height)

	return leftRigh || topBottom
}

func isVisible(trees []int, height int) bool {
	for i := 0; i < len(trees); i++ {
		if trees[i] >= height {
			return false
		}
	}
	return true
}

func reverseArray(arr []int) []int {
	reversed := make([]int, len(arr))
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		reversed[j] = arr[i]
		j++
	}
	return reversed
}
