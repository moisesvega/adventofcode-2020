package main

import (
	"log"

	"github.com/moisesvega/adventofcode-2020/utils"
)

const tree = "#"
const open = "."

// This day I won't add a test case as it will take a long time
func main() {

	//  Part 1
	log.Println(trees(3, 1))

	// Part 2
	a := trees(1, 1)
	b := trees(3, 1)
	c := trees(5, 1)
	d := trees(7, 1)
	e := trees(1, 2)

	log.Println(a, b, c, d, e)

	log.Println(a * b * c * d * e)
}

func trees(right, down int) int {
	lines := utils.ReadLinesFromFile("input.in")

	count, r, d := 0, right, down
	for d < len(lines) {
		if r >= len(lines[d]) {
			r = r - len(lines[d])
		}

		if string(lines[d][r]) == tree {
			count++
		}

		d += down
		r += right
	}

	return count
}
