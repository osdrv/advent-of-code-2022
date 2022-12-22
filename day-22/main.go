package main

import (
	"os"
)

const (
	EMPTY = 0
	OPEN  = 1
	WALL  = 2
)

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

var (
	STEPS = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

type Jumper func(map[Point2]int, Point2, int) (Point2, int)

func jumpFlat(field map[Point2]int, p Point2, dir int) (Point2, int) {
	tracep := p
	for field[tracep] != EMPTY {
		p = tracep
		tracep.x -= STEPS[dir][1]
		tracep.y -= STEPS[dir][0]
	}
	return p, dir
}

func jumpSpace50(_ map[Point2]int, p Point2, dir int) (Point2, int) {
	SIDE := 50
	x, y := p.x, p.y
	if y == 0 && x >= SIDE && x < 2*SIDE && dir == UP {
		return Point2{x: 0, y: 3*SIDE + (x - SIDE)}, RIGHT
	}
	if y == 0 && x >= 2*SIDE && x < 3*SIDE && dir == UP {
		return Point2{x: 0 + (x - 2*SIDE), y: 4*SIDE - 1}, UP
	}
	if x == 3*SIDE-1 && y >= 0 && y < SIDE && dir == RIGHT {
		return Point2{x: 2*SIDE - 1, y: 3*SIDE - 1 - y}, LEFT
	}
	if y == SIDE-1 && x >= 2*SIDE && x < 3*SIDE && dir == DOWN {
		return Point2{x: 2*SIDE - 1, y: SIDE + (x - 2*SIDE)}, LEFT
	}
	if x == 2*SIDE-1 && y >= SIDE && y < 2*SIDE && dir == RIGHT {
		return Point2{x: 2*SIDE + (y - SIDE), y: SIDE - 1}, UP
	}
	if x == 2*SIDE-1 && y >= 2*SIDE && y < 3*SIDE && dir == RIGHT {
		return Point2{x: 3*SIDE - 1, y: SIDE - 1 - (y - 2*SIDE)}, LEFT
	}
	if y == 3*SIDE-1 && x >= SIDE && x < 2*SIDE && dir == DOWN {
		return Point2{x: SIDE - 1, y: 3*SIDE + (x - SIDE)}, LEFT
	}
	if x == SIDE-1 && y >= 3*SIDE && y < 4*SIDE && dir == RIGHT {
		return Point2{x: SIDE + (y - 3*SIDE), y: 3*SIDE - 1}, UP
	}
	if y == 4*SIDE-1 && x >= 0 && x < SIDE && dir == DOWN {
		return Point2{x: 2*SIDE + x, y: 0}, DOWN
	}
	if x == 0 && y >= 3*SIDE && y < 4*SIDE && dir == LEFT {
		return Point2{x: SIDE + (y - 3*SIDE), y: 0}, DOWN
	}
	if x == 0 && y >= 2*SIDE && y < 3*SIDE && dir == LEFT {
		return Point2{x: SIDE, y: SIDE - 1 - (y - 2*SIDE)}, RIGHT
	}
	if y == 2*SIDE && x >= 0 && x < SIDE && dir == UP {
		return Point2{x: SIDE, y: SIDE + x}, RIGHT
	}
	if x == SIDE && y >= SIDE && y < 2*SIDE && dir == LEFT {
		return Point2{x: y - SIDE, y: 2 * SIDE}, DOWN
	}
	if x == SIDE && y >= 0 && y < SIDE && dir == LEFT {
		return Point2{x: 0, y: 3*SIDE - 1 - y}, RIGHT
	}
	panic("missing edge")
}

func jumpSpace4(field map[Point2]int, p Point2, dir int) (Point2, int) {
	SIDE := 4
	x, y := p.x, p.y
	if x == 2*SIDE && y >= 0 && y < SIDE && dir == LEFT {
		return Point2{x: SIDE + y, y: SIDE}, DOWN
	}
	if y == 0 && x >= 2*SIDE && x < 3*SIDE && dir == UP {
		return Point2{x: SIDE - 1 - (x - 2*SIDE), y: SIDE}, DOWN // 1 -> 2
	}
	if x == 3*SIDE-1 && y >= 0 && y < SIDE && dir == RIGHT {
		return Point2{x: 4*SIDE - 1, y: 3*SIDE - 1 - y}, LEFT
	}
	if x == 3*SIDE-1 && y >= SIDE && y < 2*SIDE && dir == RIGHT {
		return Point2{x: 4*SIDE - 1 - (y - SIDE), y: 2 * SIDE}, DOWN
	}
	if y == 2*SIDE && x >= 3*SIDE && x < 4*SIDE && dir == UP {
		return Point2{x: 3*SIDE - 1, y: 2*SIDE - 1 - (x - 3*SIDE)}, LEFT
	}
	if x == 4*SIDE-1 && y >= 2*SIDE && y < 3*SIDE && dir == RIGHT {
		return Point2{x: 3*SIDE - 1, y: SIDE - 1 - (y - 2*SIDE)}, LEFT
	}
	if y == 3*SIDE-1 && x >= 3*SIDE && x < 4*SIDE && dir == DOWN {
		return Point2{x: 0, y: 2*SIDE - 1 - (x - 3*SIDE)}, RIGHT
	}
	if y == 3*SIDE-1 && x >= 2*SIDE && x < 3*SIDE && dir == DOWN {
		return Point2{x: SIDE - 1 - (x - 2*SIDE), y: 2*SIDE - 1}, UP
	}
	if x == 2*SIDE && y >= 2*SIDE && y < 3*SIDE && dir == LEFT {
		return Point2{x: 2*SIDE - 1 - (y - 2*SIDE), y: 2*SIDE - 1}, UP
	}
	if y == 2*SIDE-1 && x >= SIDE && x < 2*SIDE-1 && dir == DOWN {
		return Point2{x: 2 * SIDE, y: 3*SIDE - 1 - (x - SIDE)}, RIGHT
	}
	if y == 2*SIDE-1 && x >= 0 && x < SIDE && dir == DOWN {
		return Point2{x: 3*SIDE - 1 - x, y: 3*SIDE - 1}, UP
	}
	if x == 0 && y >= SIDE && y < 2*SIDE && dir == LEFT {
		return Point2{x: 4*SIDE - 1 - (y - SIDE), y: 3*SIDE - 1}, UP
	}
	if y == SIDE && x >= 0 && x < SIDE && dir == UP {
		return Point2{x: 3*SIDE - 1 - x, y: 0}, DOWN
	}
	if y == SIDE && x >= SIDE && x < 2*SIDE && dir == UP {
		return Point2{x: 2 * SIDE, y: x - SIDE}, RIGHT
	}
	panic("missing edge")
}

func tracePath(field map[Point2]int, start Point2, dir int, path string, jmp Jumper) (Point2, int) {
	debugf("starting at %+v facing %d", start, dir)

	ptr := 0
	p := start
	for ptr < len(path) {
		if isNumber(path[ptr]) {
			var nsteps int
			nsteps, ptr = readInt(path, ptr)
			debugf("Moving %d steps", nsteps)
			for s := 0; s < nsteps; s++ {
				nextp := Point2{p.x + STEPS[dir][1], p.y + STEPS[dir][0]}
				nextdir := dir
				if field[nextp] == EMPTY {
					nextp, nextdir = jmp(field, p, dir)
					debugf("Jump to %+v", nextp)
				}
				if field[nextp] == WALL {
					break
				}
				p = nextp
				dir = nextdir
			}
		} else {
			turn := path[ptr]
			ptr++
			debugf("Turning %c", turn)
			if turn == 'R' {
				dir++
				dir %= 4
			} else if turn == 'L' {
				dir--
				if dir < 0 {
					dir += 4
				}
			}
		}
	}

	return p, dir
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := make(map[Point2]int)

	var start *Point2

	i := 0
	for i < len(lines) {
		if len(lines[i]) == 0 {
			i++
			break
		}
		for j := 0; j < len(lines[i]); j++ {
			p := Point2{x: j, y: i}
			if lines[i][j] == '.' {
				if start == nil {
					start = &Point2{x: j, y: i}
				}
				field[p] = OPEN
			} else if lines[i][j] == '#' {
				field[p] = WALL
			}
		}
		i++
	}
	path := lines[i]

	debugf("field: %+v", field)
	debugf("path: %s", path)

	end, dir := tracePath(field, *start, RIGHT, path, jumpFlat)

	printf("end: %+v, dir: %d", end, dir)
	printf("password: %d", 1000*(end.y+1)+4*(end.x+1)+dir)

	//end, dir = tracePath(field, *start, RIGHT, path, jumpSpace4)
	end, dir = tracePath(field, *start, RIGHT, path, jumpSpace50)
	printf("end2: %+v, dir2: %d", end, dir)
	printf("password2: %d", 1000*(end.y+1)+4*(end.x+1)+dir)
}
