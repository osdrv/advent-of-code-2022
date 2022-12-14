package main

import (
	"os"
)

const (
	_ int = iota
	SAND
	ROCK
)

func addLine(f map[Point2]int, s string) {
	points := make([]Point2, 0, 1)
	ptr := 0
	var x, y int
	for ptr < len(s) {
		x, ptr = readInt(s, ptr)
		ptr = consume(s, ptr, ',')
		y, ptr = readInt(s, ptr)
		if ptr < len(s) {
			_, ptr = readStr(s, ptr, " -> ")
		}
		points = append(points, Point2{x, y})
	}

	for i := 1; i < len(points); i++ {
		p1, p2 := points[i-1], points[i]
		step := [2]int{sign(p2.x - p1.x), sign(p2.y - p1.y)}
		p := p1
		for {
			f[p] = ROCK
			if p == p2 {
				break
			}
			p.x += step[0]
			p.y += step[1]
		}
	}
}

func sign(v int) int {
	if v == 0 {
		return 0
	} else if v < 0 {
		return -1
	}
	return 1
}

func simulate(f map[Point2]int, src Point2) int {
	maxY := -ALOT
	for p := range f {
		maxY = max(maxY, p.y)
	}

	units := 0
	for {
		units++
		p := src
		for p.y <= maxY {
			if f[Point2{p.x, p.y + 1}] == 0 {
				p.y++
			} else if f[Point2{p.x - 1, p.y + 1}] == 0 {
				p.x--
				p.y++
			} else if f[Point2{p.x + 1, p.y + 1}] == 0 {
				p.x++
				p.y++
			} else {
				f[p] = SAND
				break
			}
			if p.y >= maxY {
				return units - 1
			}
		}
	}
}

func simulate2(f map[Point2]int, src Point2) int {
	maxY := -ALOT
	for p := range f {
		maxY = max(maxY, p.y)
	}

	maxY += 2

	bottom := func(p Point2) int {
		if p.y >= maxY {
			return ROCK
		}
		return f[p]
	}

	units := 0
	for {
		units++
		p := src
		for p.y < maxY {
			if bottom(Point2{p.x, p.y + 1}) == 0 {
				p.y++
			} else if bottom(Point2{p.x - 1, p.y + 1}) == 0 {
				p.x--
				p.y++
			} else if bottom(Point2{p.x + 1, p.y + 1}) == 0 {
				p.x++
				p.y++
			} else {
				if p == src {
					return units
				}
				f[p] = SAND
				break
			}
			debugf("unit %d is at point %+v", units, p)
		}
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := make(map[Point2]int)

	for _, line := range lines {
		addLine(field, line)
	}

	debugf("field: %+v (%d)", field, len(field))

	src := Point2{500, 0}

	//printf("units: %d", simulate2(field, src))
	printf("units: %d", simulate2(field, src))
}
