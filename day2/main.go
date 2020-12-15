package main

import (
	"strconv"
	"strings"
)

type request struct {
	password string
	target   string
	min      int
	max      int
}

func valid(in *request) (valid bool) {
	count := strings.Count(in.password, in.target)
	return count >= in.min && count <= in.max
}

func pt1(in []*request) (out int) {
	out = 0
	for _, r := range in {
		if valid(r) {
			out++
		}
	}
	return out
}

// 3-6 r: frrhxsnrmgmw
func parseLine(in string) (out *request) {
	out = &request{}
	split := strings.Split(in, " ")

	// first part min-max
	minMax := strings.Split(split[0], "-")
	out.min, _ = strconv.Atoi(minMax[0])
	out.max, _ = strconv.Atoi(minMax[1])

	// target
	target := strings.Split(split[1], ":")
	out.target = target[0]

	// password
	out.password = split[2]

	return out
}

func parseLines(in []string) (out []*request) {
	out = make([]*request, 0)
	for _, s := range in {
		out = append(out, parseLine(s))
	}
	return out
}
