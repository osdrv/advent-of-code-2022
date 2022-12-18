package main

import (
	"os"
)

var (
	STEPS = [][3]int{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	pmap := make(map[Point3]bool)

	xmin, xmax, ymin, ymax, zmin, zmax := ALOT, -ALOT, ALOT, -ALOT, ALOT, -ALOT
	for _, line := range lines {
		nn := parseInts(line)
		p := Point3{nn[0], nn[1], nn[2]}
		xmin = min(xmin, p.x)
		xmax = max(xmax, p.x)
		ymin = min(ymin, p.y)
		ymax = max(ymax, p.y)
		zmin = min(zmin, p.z)
		zmax = max(zmax, p.z)
		pmap[p] = true
	}

	sides := make(map[[6]int]bool)
	for p := range pmap {
		for _, s := range STEPS {
			p1 := Point3{p.x + s[0], p.y + s[1], p.z + s[2]}
			sd := [6]int{p.x, p.y, p.z, p1.x, p1.y, p1.z}
			if _, ok := pmap[p1]; !ok {
				sides[sd] = true
			}
		}
	}

	printf("exposed sides: %d", len(sides))

	var exploreAirGap func(Point3, map[Point3]bool) bool
	exploreAirGap = func(p Point3, visited map[Point3]bool) bool {
		if _, ok := visited[p]; ok {
			return false
		}
		visited[p] = true
		for _, s := range STEPS {
			p1 := Point3{p.x + s[0], p.y + s[1], p.z + s[2]}
			if _, ok := pmap[p1]; ok {
				continue
			}
			if p1.x < xmin || p1.x > xmax || p1.y < ymin || p1.y > ymax || p1.z < zmin || p1.z > zmax {
				return true
			}
			if exploreAirGap(p1, visited) {
				return true
			}
		}
		return false
	}

	exterior := make(map[[6]int]bool)

	for side := range sides {
		if inf := exploreAirGap(Point3{side[3], side[4], side[5]}, make(map[Point3]bool)); inf {
			exterior[side] = true
		}
	}

	printf("exterior sides: %d", len(exterior))
}
