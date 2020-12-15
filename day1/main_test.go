package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moisesvega/adventofcode-2020/utils"
)

func TestDay1(t *testing.T) {
	tt := map[string]func(param *[]int) (expectedResponse int){
		"success": func(param *[]int) (expectedResponse int) {
			in := &[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			}

			*param = append(*param, *in...)

			expectedResponse = 514579
			return
		},
		"from input file": func(param *[]int) (expectedResponse int) {
			file, err := os.Open("input.in")
			if err != nil {
				panic(err.Error())
			}
			in := &[]int{}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				value, err := strconv.Atoi(scanner.Text())
				if err != nil {
					panic(err.Error())
				}
				*in = append(*in, value)
			}

			*param = append(*param, *in...)
			expectedResponse = 618144
			return
		},
	}

	for name, f := range tt {
		f := f
		t.Run(name, func(t *testing.T) {
			in := &[]int{}
			expectedResponse := f(in)
			response := day1(*in, 2020)
			assert.Equal(t, expectedResponse, response)
		})
	}
}

func TestDay1Pt2(t *testing.T) {
	tt := map[string]func(param *[]int) (expectedResponse int){
		"success": func(param *[]int) (expectedResponse int) {
			in := &[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			}

			*param = append(*param, *in...)

			expectedResponse = 241861950
			return
		},
		"from input file": func(param *[]int) (expectedResponse int) {
			in := &[]int{}
			lines := utils.ReadLinesFromFile("input.in")
			for _, line := range lines {
				value, err := strconv.Atoi(line)
				if err != nil {
					panic(err.Error())
				}
				*in = append(*in, value)
			}

			*param = append(*param, *in...)
			expectedResponse = 173538720
			return
		},
	}

	for name, f := range tt {
		f := f
		t.Run(name, func(t *testing.T) {
			in := &[]int{}
			expectedResponse := f(in)
			response := day1Pt2(*in, 2020)
			assert.Equal(t, expectedResponse, response)
		})
	}
}
