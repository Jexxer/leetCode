package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if !strings.Contains(haystack, needle) {
		return -1
	}
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
