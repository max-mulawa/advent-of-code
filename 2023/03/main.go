package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type part struct {
	value int
	rowNo int
	i     int
	len   int
	valid bool
}

func (p *part) validate(rows []string) {
	for _, r := range rows {
		for cIdx, c := range r {
			if cIdx >= (p.i-1) && cIdx <= (p.i+p.len) && isSymbol(c) {
				p.valid = true
				return
			}
		}
	}
}

func isSymbol(c rune) bool {
	return c != '.' && !unicode.IsDigit(c)
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

	s := bufio.NewScanner(f)
	total := 0
	rowCount := 0
	var parts []*part
	var rows []string
	var p *part
	for s.Scan() {
		row := s.Text()
		rows = append(rows, row)
	}

	for rowIdx, row := range rows {
		var rowsToInspect []string
		if rowIdx-1 >= 0 {
			rowsToInspect = append(rowsToInspect, rows[rowIdx-1])
		}
		rowsToInspect = append(rowsToInspect, row)
		if rowIdx+1 < len(rows) {
			rowsToInspect = append(rowsToInspect, rows[rowIdx+1])
		}

		partTokenStart := -1
		for i, c := range row {
			if unicode.IsDigit(c) && partTokenStart == -1 {
				if p == nil {
					p = &part{
						rowNo: rowCount,
						i:     i,
					}
					parts = append(parts, p)
					partTokenStart = i

				}
			}

			if !unicode.IsDigit(c) || i == (len(row)-1) {
				if partTokenStart != -1 {
					p.value, _ = strconv.Atoi(row[p.i:i])
					p.len = i - p.i
					p.validate(rowsToInspect)

					p = nil
					partTokenStart = -1
				}
			}
		}

		rowCount++
	}

	for _, p := range parts {
		if p.valid {
			fmt.Println(*p)
			total += p.value
		}
	}

	fmt.Println(total)

}
