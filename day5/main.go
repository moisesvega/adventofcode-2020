package main

import (
	"log"
	"math"
	"sort"

	"github.com/moisesvega/adventofcode-2020/utils"
)

func main() {

	lines := utils.ReadLinesFromFile("input.in")

	max := 0
	ids := make([]int, 0)
	for _, line := range lines {
		id := calculateID(line)

		if id > max {
			max = id
		}
		// log.Println(id)
		ids = append(ids, id)
	}
	// log.Println(max)

	// Part 2
	sort.Ints(ids)
	for i, id := range ids {
		if i == len(ids)-1 {
			continue
		}
		if id+1 != ids[i+1] {
			log.Println(id, ids[i+1])
		}
	}

	// log.Println(ids)
}

// Start by considering the whole range, rows 0 through 127.
// F means to take the lower half, keeping rows 0 through 63.
// B means to take the upper half, keeping rows 32 through 63.
// F means to take the lower half, keeping rows 32 through 47.
// B means to take the upper half, keeping rows 40 through 47.
// B keeps rows 44 through 47.
// F keeps rows 44 through 45.
// The final F keeps the lower of the two, row 44.

func calculateID(in string) (out int) {

	first, last := 0, 127
	columnR, columnL := 0, 7
	out = 0
	column := 0
	for _, c := range in {
		calr := math.Ceil((float64(last) - float64(first)) / 2)
		calc := math.Ceil((float64(columnL) - float64(columnR)) / 2)

		switch string(c) {
		case "F":
			last = last - int(calr)
			out = first
		case "B":
			first = first + int(calr)
			out = last
		case "L":
			columnL = columnL - int(calc)
			column = columnR
		case "R":
			columnR = columnR + int(calc)
			column = columnL
		}
		// log.Println(first, last, columnR, columnL)
	}
	// log.Println(out, column, columnR, columnL)
	out = out*8 + column

	return out
}
