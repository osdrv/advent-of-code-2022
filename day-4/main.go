package main

import (
	"os"
	"strings"
)

func parseRange(s string) [2]int {
	chs := strings.SplitN(s, "-", 2)
	return [2]int{parseInt(chs[0]), parseInt(chs[1])}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	contained := 0
	overlaps := 0

	pairs := make([][2][2]int, 0, len(lines))
	for _, line := range lines {
		chs := strings.SplitN(line, ",", 2)
		r1, r2 := parseRange(chs[0]), parseRange(chs[1])
		pairs = append(pairs, [2][2]int{r1, r2})

		if isContainedWithin(r1, r2) || isContainedWithin(r2, r1) {
			contained++
		}

		if isOverlapsWith(r1, r2) {
			debugf("%+v overlaps with %+v", r1, r2)
			overlaps++
		}
	}
	printf("contained: %d", contained)
	printf("overlaps: %d", overlaps)
}

func isContainedWithin(r1, r2 [2]int) bool {
	return r2[0] <= r1[0] && r2[1] >= r1[1]
}

func isOverlapsWith(r1, r2 [2]int) bool {
	return (r1[0] >= r2[0] && r1[0] <= r2[1]) || (r1[1] >= r2[0] && r1[1] <= r2[1]) || (r2[0] >= r1[0] && r2[0] <= r1[1]) || (r2[1] >= r1[0] && r2[1] <= r1[1])
}
