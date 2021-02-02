package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := map[string]int{}
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	return m
}

func main() {
	wc.Test(wordCount)
}
