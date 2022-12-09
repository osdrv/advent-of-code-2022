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
	steps = map[int][2]int{
		UP:    {0, -1},
		RIGHT: {1, 0},
		DOWN:  {0, 1},
		LEFT:  {-1, 0},
	}
)

func parseMove(s string) [2]int {
	var move [2]int
	switch s[0] {
	case 'U':
		move[0] = UP
	case 'R':
		move[0] = RIGHT
	case 'D':
		move[0] = DOWN
	case 'L':
		move[0] = LEFT
	}
	move[1] = parseInt(s[2:])

	return move
}

func sign(v int) int {
	if v == 0 {
		return 0
	} else if v > 0 {
		return 1
	}
	return -1
}

func isTouch(p1, p2 Point2) bool {
	return abs(p1.x-p2.x) <= 1 && abs(p1.y-p2.y) <= 1
}

func makeRope(knots int) []Point2 {
	rope := make([]Point2, 0, knots)
	for i := 0; i < knots; i++ {
		rope = append(rope, Point2{0, 0})
	}
	return rope
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	moves := make([][2]int, 0, len(lines))

	for _, line := range lines {
		moves = append(moves, parseMove(line))
	}

	visited := make(map[Point2]bool)

	rope := makeRope(10)

	for _, move := range moves {
		debugf("move: %+v", move)
		for i := 0; i < move[1]; i++ {
			step := steps[move[0]]
			rope[0].x += step[0]
			rope[0].y += step[1]
			debugf("head moves to %+v", rope[0])

			for i := 1; i < len(rope); i++ {
				if !isTouch(rope[i], rope[i-1]) {
					rope[i].x += sign(rope[i-1].x - rope[i].x)
					rope[i].y += sign(rope[i-1].y - rope[i].y)
				}
			}
			visited[rope[len(rope)-1]] = true
		}
	}

	debugf("visited: %+v", visited)
	printf("visited count: %d", len(visited))
}
