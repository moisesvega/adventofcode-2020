package main

import (
	"log"

	"github.com/moisesvega/adventofcode-2020/utils"
)

func main() {

	groupsOfGroups := getGroups()
	count := 0
	for _, group := range groupsOfGroups {
		// count += uniques(group)
		count += uniqueInGroup(group)
	}

	log.Println(count)
}

func uniques(group []string) (count int) {
	m := map[string]bool{}
	for _, s := range group {
		for _, c := range s {
			if _, ok := m[string(c)]; !ok {
				m[string(c)] = true
			}
		}
	}
	log.Println(len(m), m, group)
	return len(m)
}

func uniqueInGroup(group []string) (count int) {
	m := map[string]int{}
	max := 0
	for _, s := range group {
		for _, c := range s {
			m[string(c)] += 1
			if m[string(c)] == len(group) {
				max += 1
			}
		}
	}
	// log.Println(len(m), m, group, max)
	return max
}

func getGroups() [][]string {
	lines := utils.ReadLinesFromFile("input.in")

	groupsOfGroups := make([][]string, 0)
	group := make([]string, 0)

	for _, line := range lines {
		if len(line) == 0 {
			groupsOfGroups = append(groupsOfGroups, group)
			group = make([]string, 0)
			continue
		}
		group = append(group, line)
	}
	groupsOfGroups = append(groupsOfGroups, group)
	return groupsOfGroups
}
