package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moisesvega/adventofcode-2020/utils"
)

func TestValid(t *testing.T) {
	t.Parallel()

	tests := map[string]func() (rq *request, expectedResult bool){
		"valid": func() (rq *request, expectedResult bool) {
			rq = &request{
				password: "abcde",
				target:   "a",
				min:      1,
				max:      3,
			}
			expectedResult = true

			return
		},
		"invalid": func() (rq *request, expectedResult bool) {
			rq = &request{
				password: "cdefg",
				target:   "b",
				min:      1,
				max:      3,
			}
			expectedResult = false

			return
		},
	}

	for name, f := range tests {
		f := f
		t.Run(name, func(t *testing.T) {
			r, expectedResult := f()

			result := valid(r)
			assert.Equal(t, expectedResult, result)
		})
	}
}

func TestPt1(t *testing.T) {
	t.Parallel()

	tests := map[string]func() (rq []*request, expectedResult int){
		"valid": func() (rq []*request, expectedResult int) {
			in := []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			}

			rq = parseLines(in)
			expectedResult = 2
			return
		},
		"valid from file": func() (rq []*request, expectedResult int) {
			in := utils.ReadLinesFromFile("input.in")

			rq = parseLines(in)
			expectedResult = 628
			return
		},
	}

	for name, f := range tests {
		f := f
		t.Run(name, func(t *testing.T) {
			r, expectedResult := f()

			result := pt1(r)
			assert.Equal(t, expectedResult, result)
		})
	}
}

func TestPt2(t *testing.T) {
	t.Parallel()

	tests := map[string]func() (rq []*request, expectedResult int){
		"valid": func() (rq []*request, expectedResult int) {
			in := []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			}

			rq = parseLines(in)
			expectedResult = 1
			return
		},
		"valid from file": func() (rq []*request, expectedResult int) {
			in := utils.ReadLinesFromFile("input.in")

			rq = parseLines(in)
			expectedResult = 705
			return
		},
	}

	for name, f := range tests {
		f := f
		t.Run(name, func(t *testing.T) {
			r, expectedResult := f()

			result := pt2(r)
			assert.Equal(t, expectedResult, result)
		})
	}
}
