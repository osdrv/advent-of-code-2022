package main

import (
	"os"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

var (
	STEPS = [4][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

func tracePath(maxx, maxy int, start, finish Point2, bls []Point3, rts int) int {

	if rts == 0 {
		return 0
	}

	opts := []Point2{start}

	T := 0
	for {
		T++
		debugf("=== Minute %d ===", T)
		bmap := make(map[Point2]struct{})
		for bix := 0; bix < len(bls); bix++ {
			b := bls[bix]
			b.x += STEPS[b.z][0]
			b.y += STEPS[b.z][1]
			if b.x == 0 {
				b.x = maxx - 1
			}
			if b.y == 0 {
				b.y = maxy - 1
			}
			if b.x == maxx {
				b.x = 1
			}
			if b.y == maxy {
				b.y = 1
			}
			bls[bix] = b
			bmap[Point2{b.x, b.y}] = struct{}{}
		}

		newOpts := make([]Point2, 0, 1)
		for _, opt := range opts {
			// wait
			if _, ok := bmap[opt]; !ok {
				newOpts = append(newOpts, opt)
				bmap[opt] = struct{}{}
			}
			// move
			for _, step := range STEPS {
				nopt := Point2{opt.x + step[0], opt.y + step[1]}
				if nopt == finish {
					return T + tracePath(maxx, maxy, finish, start, bls, rts-1)
				}
				if nopt.x == 0 || nopt.x == maxx || nopt.y == 0 || nopt.y == maxy {
					continue
				}
				if _, ok := bmap[nopt]; !ok {
					newOpts = append(newOpts, nopt)
					bmap[nopt] = struct{}{}
				}
			}
		}

		if len(newOpts) == 0 {
			break
		}

		opts = newOpts
	}

	return -1
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	bls := make([]Point3, 0, 1)

	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '^':
				bls = append(bls, Point3{x, y, UP})
			case '>':
				bls = append(bls, Point3{x, y, RIGHT})
			case 'v':
				bls = append(bls, Point3{x, y, DOWN})
			case '<':
				bls = append(bls, Point3{x, y, LEFT})
			}
		}
	}

	maxx, maxy := len(lines[0])-1, len(lines)-1
	bls2 := make([]Point3, len(bls))
	copy(bls2, bls)

	start := Point2{1, 0}
	finish := Point2{maxx - 1, maxy}

	debugf("bls: %v, start: %+v, finish: %+v", bls, start, finish)

	t := tracePath(maxx, maxy, start, finish, bls, 1)
	printf("min time: %d", t)

	t2 := tracePath(maxx, maxy, start, finish, bls2, 3)
	printf("min time2: %d", t2)
}
