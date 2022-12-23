package main

import (
	"bytes"
	"os"
)

const (
	ELF = 1
)

const (
	N int = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var (
	STEPS = [8][2]int{
		{-1, 0},  // N
		{-1, 1},  //NE
		{0, 1},   //E
		{1, 1},   //SE
		{1, 0},   //S
		{1, -1},  //SW
		{0, -1},  //W
		{-1, -1}, //NW
	}
)

var (
	PROPOSALS = [4][4]int{
		{N, NE, NW, N},
		{S, SE, SW, S},
		{W, NW, SW, W},
		{E, NE, SE, E},
	}
)

var (
	RETRACTED = Point2{ALOT, ALOT}
)

func minmax(f map[Point2]int) (int, int, int, int) {
	var minx, maxx, miny, maxy int

	for p := range f {
		minx = min(minx, p.x)
		maxx = max(maxx, p.x)
		miny = min(miny, p.y)
		maxy = max(maxy, p.y)
	}

	return minx, maxx, miny, maxy
}

func printField(f map[Point2]int) string {
	minx, maxx, miny, maxy := minmax(f)
	var b bytes.Buffer
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if f[Point2{x, y}] == ELF {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func evolve(f map[Point2]int, round int) (map[Point2]int, int) {
	res := make(map[Point2]int)

	props := make(map[Point2]Point2)

	debugf("field: %+v", f)

	moves := 0

	for p := range f {
		var fnd [8]bool
		fndcnt := 0
		for dir, step := range STEPS {
			np := Point2{x: p.x + step[1], y: p.y + step[0]}
			if v, ok := f[np]; ok && v == ELF {
				fnd[dir] = true
				fndcnt++
			}
		}
		if fndcnt == 0 {
			res[p] = ELF
			continue
		}

		var prop Point2
		didProp := false
		for i := 0; i < len(PROPOSALS); i++ {
			pr := PROPOSALS[(i+round)%len(PROPOSALS)]
			if !fnd[pr[0]] && !fnd[pr[1]] && !fnd[pr[2]] {
				prop = Point2{x: p.x + STEPS[pr[3]][1], y: p.y + STEPS[pr[3]][0]}
				didProp = true
				break
			}
		}
		if !didProp {
			res[p] = ELF
			continue
		}

		if prev, ok := props[prop]; ok {
			if prev != RETRACTED {
				res[prev] = ELF
			}
			res[p] = ELF
			props[prop] = RETRACTED
			continue
		}

		props[prop] = p
	}

	for prop, from := range props {
		if from != RETRACTED {
			moves++
			res[prop] = ELF
		}
	}

	return res, moves
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	field := make(map[Point2]int)

	for i := range lines {
		for j, ch := range lines[i] {
			if ch == '#' {
				field[Point2{x: j, y: i}] = ELF
			}
		}
	}

	debugf("field: %+v", field)
	print(printField(field))

	CHECKPOINT := 10

	for round := 0; round < CHECKPOINT; round++ {
		printf("=== round %d ===", round)
		field, _ = evolve(field, round)
		print(printField(field))
	}

	minx, maxx, miny, maxy := minmax(field)
	empty := (maxx-minx+1)*(maxy-miny+1) - len(field)
	printf("empty tiles: %d", empty)

	round := CHECKPOINT
	moves := 0
	for {
		field, moves = evolve(field, round)
		if DEBUG {
			printf("=== round %d ===", round+1)
			println(printField(field))
		}
		round++
		if moves == 0 {
			printf("no-move round: %d", round)
			break
		}
	}
}
