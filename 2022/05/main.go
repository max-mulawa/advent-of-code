package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

//1,5,9,13,17,.....

type createsConfig struct {
	columns []stack.Stack
}

type move struct {
	src int
	dst int
	cnt int
}

func main() {
	var columnCnt = 9
	cratesPayload, err := os.ReadFile("input-crates.txt")
	// var columnCnt = 3
	// cratesPayload, err := os.ReadFile("input-crates-test.txt")
	if err != nil {
		panic(err)
	}
	movesPayload, err := os.ReadFile("input-moves.txt")
	// movesPayload, err := os.ReadFile("input-moves-test.txt")

	cratesCfg := &createsConfig{
		columns: make([]stack.Stack, columnCnt),
	}

	cratesCfgLines := []string{}
	r := bufio.NewReader(bytes.NewReader(cratesPayload))
	for {
		crates, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		tmp := strings.TrimSpace(crates)
		if tmp == "" {
			break
		}

		cratesCfgLines = append(cratesCfgLines, crates)

	}

	for i := len(cratesCfgLines) - 1; i >= 0; i-- {
		idx := 1
		for cNo := 0; cNo < columnCnt; cNo++ {
			crate := string([]byte{cratesCfgLines[i][idx]})
			if crate != " " {
				cratesCfg.columns[cNo].Push(crate)
			}
			idx += 4
		}
	}

	for _, col := range cratesCfg.columns {
		fmt.Printf("%s", col.Peek().(string))
	}
	fmt.Println()

	// parse moves
	moves := []move{}

	if err != nil {
		panic(err)
	}

	r = bufio.NewReader(bytes.NewReader(movesPayload))
	for {
		movesTxt, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		movesTxt = strings.TrimSpace(movesTxt)
		if movesTxt == "" {
			break
		}

		movesTokens := strings.Split(movesTxt, " ")

		cnt, _ := strconv.Atoi(movesTokens[1])
		src, _ := strconv.Atoi(movesTokens[3])
		dst, _ := strconv.Atoi(movesTokens[5])

		mv := move{
			src: src,
			dst: dst,
			cnt: cnt,
		}
		moves = append(moves, mv)
	}

	fmt.Println(moves)

	for _, m := range moves {
		executeMoveV2(cratesCfg, m)
	}

	for _, col := range cratesCfg.columns {
		if col.Len() > 0 {
			c := col.Peek().(string)
			fmt.Printf("%s", c)
		} else {
			fmt.Printf(" ")
		}

	}
}

func executeMoveV2(cratesCfg *createsConfig, m move) {
	crates := []string{}
	for i := 0; i < m.cnt; i++ {
		c := cratesCfg.columns[m.src-1].Pop()
		crates = append(crates, c.(string))
	}
	for i := m.cnt - 1; i >= 0; i-- {
		cratesCfg.columns[m.dst-1].Push(crates[i])
	}
}

func executeMove(cratesCfg *createsConfig, m move) {
	for i := 0; i < m.cnt; i++ {
		c := cratesCfg.columns[m.src-1].Pop()
		cratesCfg.columns[m.dst-1].Push(c)
	}
}
