package main

import (
	"os"
)

func parseScan(s string) (p1, p2 Point2) {
	var xs, ys, xb, yb int
	Scanf(s, "Sensor at x={int}, y={int}: closest beacon is at x={int}, y={int}", &xs, &ys, &xb, &yb)
	return Point2{xs, ys}, Point2{xb, yb}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	sensors := make([][3]int, 0, 1)

	ix := 0
	y0 := parseInt(lines[ix])
	ix++
	maxy := parseInt(lines[ix])
	ix += 2

	nox := make(map[int]bool)
	for ix < len(lines) {
		s, b := parseScan(lines[ix])
		dist := s.Dist(b)
		sensors = append(sensors, [3]int{s.x, s.y, dist})
		debugf("s: %+v, b: %+v, d: %d", s, b, dist)

		if s.y == y0 {
			nox[s.x] = true
		}
		if b.y == y0 {
			nox[b.x] = true
		}

		ix++
	}

	cnt := 0
	ranges := intersect(y0, sensors)
	debugf("ranges: %+v", ranges)
	for _, rr := range ranges {
		cnt += (1 + rr[1] - rr[0])
		for nx := range nox {
			if nx >= rr[0] && nx <= rr[1] {
				cnt--
			}
		}
	}

	printf("cnt: %d", cnt)

	for y := 0; y <= maxy; y++ {
		rrs := intersect(y, sensors)
		if len(rrs) == 1 {
			continue
		}
		if rrs[1][0]-rrs[0][1] <= 1 {
			continue
		}
		printf("vacant point: y=%d, x=%d", y, rrs[0][1]+1)
		freq := uint64(rrs[0][1]+1)*4000000 + uint64(y)
		printf("freq: %d", freq)
		break
	}

}

func intersect(y0 int, sensors [][3]int) []Range {
	ranges := make([]Range, 0, 1)
	for _, s := range sensors {
		x, y, d := s[0], s[1], s[2]
		if y-d <= y0 && y+d >= y0 {
			dy := abs(y - y0)
			ranges = append(ranges, NewRange(x-(d-dy), x+(d-dy)))
		}
	}

	return mergeRanges(ranges)
}
