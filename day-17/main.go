package main

import (
	"os"
)

var (
	SHAPES = [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},             // -
		{{0, 1}, {-1, 0}, {-1, 1}, {-1, 2}, {-2, 1}}, // +
		{{0, 0}, {0, 1}, {0, 2}, {-1, 2}, {-2, 2}},   // inversed L
		{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},          // |
		{{0, 0}, {0, 1}, {-1, 0}, {-1, 1}},           // .
	}
)

func draw(f map[Point2]bool, shape [][2]int, x, y int) {
	for _, p := range shape {
		pp := Point2{y: y + p[0], x: x + p[1]}
		f[pp] = true
	}
}

func erase(f map[Point2]bool, shape [][2]int, x, y int) {
	for _, p := range shape {
		pp := Point2{y: y + p[0], x: x + p[1]}
		delete(f, pp)
	}
}

func tryMove(f map[Point2]bool, shape [][2]int, x, y int, mv [2]int) bool {
	for _, p := range shape {
		if p[1]+x+mv[1] < 0 || p[1]+x+mv[1] >= 7 {
			return false
		}
		if p[0]+y+mv[0] > 0 {
			return false
		}
	}
	erase(f, shape, x, y)
	for _, p := range shape {
		pp := Point2{y: y + p[0] + mv[0], x: x + p[1] + mv[1]}
		if f[pp] {
			draw(f, shape, x, y)
			return false
		}
	}
	draw(f, shape, x+mv[1], y+mv[0])
	return true
}

func findDims(f map[Point2]bool) (int, int, int, int) {
	xmin, xmax, ymin, ymax := 0, 6, ALOT, 0
	for p := range f {
		xmin = min(xmin, p.x)
		xmax = max(xmax, p.x)
		ymin = min(ymin, p.y)
		ymax = max(ymax, p.y)
	}
	return xmin, xmax, ymin, ymax
}

func drawField(f map[Point2]bool) string {
	xmin, xmax, ymin, ymax := findDims(f)
	debugf("%d %d %d %d ", xmin, xmax, ymin, ymax)
	ff := makeByteField(ymax-ymin+1, xmax-xmin+1)
	for p := range f {
		ff[p.y-ymin][p.x-xmin] = 1
	}
	return printNumFieldWithSubs(ff, "", map[byte]string{
		0: ".",
		1: "@",
	})
}

func fieldToNums(f map[Point2]bool, height int) []byte {
	nums := make([]byte, height)
	for p := range f {
		nums[abs(p.y)] |= 1 << p.x
	}
	return nums
}

var (
	MOVES = map[byte]int{
		'<': -1,
		'>': 1,
	}
)

func play(moves string, shapes int) (map[Point2]bool, []int, int) {
	mvptr := 0

	field := make(map[Point2]bool)
	miny := 1

	history := make([]int, shapes)

	for i := 0; i < shapes; i++ {
		shape := SHAPES[i%len(SHAPES)]
		x, y := 2, miny-3-1

		draw(field, shape, x, y)

		for {
			if tryMove(field, shape, x, y, [2]int{0, MOVES[moves[mvptr]]}) {
				x += MOVES[moves[mvptr]]
			}
			mvptr++
			mvptr %= len(moves)

			if tryMove(field, shape, x, y, [2]int{1, 0}) {
				y += 1
			} else {
				break
			}
		}
		_, _, miny, _ = findDims(field)
		history[i] = 1 - miny
	}

	return field, history, abs(miny) + 1
}

// head, pattern
func findRepPattern(s []byte) ([]byte, []byte) {
	for off := 0; off < len(s); off++ {
		for patLen := 10; patLen <= (len(s)-off)/2; patLen++ {
			if eql(s[off:off+patLen], s[off+patLen:off+2*patLen]) {
				return s[:off], s[off : off+patLen]
			}
		}
	}
	return nil, nil
}

func eql(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field, history, h := play(lines[0], 2022)
	printf("Part1: %d", h)

	/*
		Part2 Thoughts

		There is a pattern in numbers

		Total number of LINES = HEAD + N * (PERIOD) + TAIL

		We can compute HEAD and PERIOD: eye-balling indicates the test input comes as:
		25 + N * (54) + TAIL

		Shape-wise it would be similar:
		SHAPES = HEAD_SHAPES + N * (SHAPE_PERIOD) + SHAPE_TAIL

	*/

	// we need to process some more shapes
	field, history, h = play(lines[0], 4000)
	nums := fieldToNums(field, h)
	debugf("nums: %+v", nums)
	head, pat := findRepPattern(nums)
	debugf("history: %+v", history)
	debugf("head: %d, pat: %d", len(head), len(pat))
	debugf("head: %+v", head)
	debugf("pat: %+v", pat)

	assert(len(head) > 0 && len(pat) > 0, "pattern could not be found, increase the number of shapes")

	hix, pix := make([]int, 0, 1), make([]int, 0, 1)
	for sh, ln := range history {
		if ln == len(head) {
			hix = append(hix, sh)
		}
		if ln == len(head)+len(pat) {
			pix = append(pix, sh)
		}
		if ln > len(head)+len(pat) {
			break
		}
	}

	debugf("hix: %+v, pix: %+v", hix, pix)

	//NSHAPES = 2022
	NSHAPES := 1000000000000

	lls := make([]uint64, 0, 1)
	for _, h := range hix {
		for _, p := range pix {
			sHead := h + 1
			sPat := p - sHead + 1
			debugf("sHead: %d, sPat: %d", sHead, sPat)

			patsl := make([]int, len(pat)+1)
			for i := 1; i <= len(pat); i++ {
				patsl[i] = history[sHead+i] - history[sHead]
			}

			lls = append(lls, uint64(len(head))+uint64((NSHAPES-sHead)/sPat)*uint64(len(pat))+uint64(patsl[(NSHAPES-sHead)%sPat]))
		}
	}

	printf("Part2: %d", lls[len(lls)-1])

}
