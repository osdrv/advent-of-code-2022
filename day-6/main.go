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

func findStartFast(s string, k int) int {
	var sa [26]int
	cnt := 0
	ix := 0
	for ix < k && ix < len(s) {
		off := int(s[ix] - 'a')
		sa[off]++
		if sa[off] == 1 {
			cnt++
		}
		ix++
	}
	for ix < len(s) && cnt != k {
		off1 := int(s[ix-k] - 'a')
		off2 := int(s[ix] - 'a')
		sa[off1]--
		if sa[off1] == 0 {
			cnt--
		}
		sa[off2]++
		if sa[off2] == 1 {
			cnt++
		}
		ix++
	}
	return ix
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		pkgstrt := findStartFast(line, 4)
		msgstrt := findStartFast(line, 14)

		printf("pkg start: %d", pkgstrt)
		printf("msg start: %d", msgstrt)
	}
}
