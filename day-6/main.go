package main

import (
	"os"
)

func findStart(s string, k int) int {
	for i := 0; i < len(s)-(k-1); i++ {
		m := make(map[byte]struct{}, k)
		for j := i; j < i+k; j++ {
			m[s[j]] = struct{}{}
		}
		if len(m) == k {
			return i + k
		}
	}
	return -ALOT
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		pkgstrt := findStart(line, 4)
		msgstrt := findStart(line, 14)
		printf("pkg start: %d", pkgstrt)
		printf("msg start: %d", msgstrt)
	}
}
