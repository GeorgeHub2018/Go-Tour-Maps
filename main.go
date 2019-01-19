package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	m := map[string]int{}
	f := strings.Fields(s)
	for _, i := range f {
		_, ok := m[i]
		if ok {
			m[i] = m[i] + 1
		} else {
			m[i] = 1
		}
	}

	return m
}

func main() {
	wc.Test(wordCount)
}
