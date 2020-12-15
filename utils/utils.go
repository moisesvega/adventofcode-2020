package utils

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(path string) (out []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	out = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}
