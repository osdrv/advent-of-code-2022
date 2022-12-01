package main

import (
	"os"
	"sort"
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	cc := make([]int, 0, 1)
	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			cc = append(cc, sum)
			sum = 0
			continue
		}
		sum += parseInt(line)
	}
	cc = append(cc, sum)

	sort.Sort(sort.Reverse(sort.IntSlice(cc)))

	printf("max sum: %d", cc[0])
	printf("top 3 sums: %d", cc[0]+cc[1]+cc[2])

}
