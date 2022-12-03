package main

import (
	"os"
)

func priorityOf(b byte) int {
	if b >= 'a' && b <= 'z' {
		return 1 + int(b-'a')
	}
	return 27 + int(b-'A')
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	sum := 0
Line:
	for _, line := range lines {
		left, right := line[:len(line)/2], line[len(line)/2:]

		printf("left: %s, right: %s", left, right)

		lm := make(map[byte]int)
		for i := 0; i < len(left); i++ {
			lm[left[i]]++
		}

		for i := 0; i < len(right); i++ {
			if _, ok := lm[right[i]]; ok {
				printf("common char: %c", right[i])
				sum += priorityOf(right[i])
				continue Line
			}
		}
	}

	printf("sum: %d", sum)

	sum2 := 0
Lines2:
	for i := 0; i < len(lines); i += 3 {
		mi, mj, mk := indexLine(lines[i]), indexLine(lines[i+1]), indexLine(lines[i+2])
		mm := mergeMaps(mi, mj, mk)

		for k, v := range mm {
			if v == 3 {
				printf("found common component: %c", k)
				sum2 += priorityOf(k)
				continue Lines2
			}
		}
	}

	printf("sum2: %d", sum2)
}

func mergeMaps(ms ...map[byte]int) map[byte]int {
	mm := make(map[byte]int)
	for _, m := range ms {
		for k, v := range m {
			mm[k] += v
		}
	}
	return mm
}

func indexLine(s string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = 1
	}
	return m
}
