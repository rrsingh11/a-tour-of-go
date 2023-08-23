package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)
	for _, val := range f {
		m[val]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
