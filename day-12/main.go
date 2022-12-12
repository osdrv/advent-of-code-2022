package main

import (
	"os"
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	var start, end Point2

	field := make([][]byte, len(lines))
	for i, line := range lines {
		field[i] = []byte(line)
		for j := 0; j < len(line); j++ {
			if field[i][j] == 'S' {
				start = Point2{i, j}
			} else if field[i][j] == 'E' {
				end = Point2{i, j}
			}
		}
	}

	minSteps := findShortestPath(field, start, end)
	println(minSteps)

	allMinSteps := ALOT
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field); j++ {
			if getHeight(field, i, j) == 0 {
				allMinSteps = min(allMinSteps, findShortestPath(field, Point2{i, j}, end))
			}
		}
	}
	println(allMinSteps)
}

func getHeight(field [][]byte, i, j int) int {
	if field[i][j] == 'S' {
		return 0
	} else if field[i][j] == 'E' {
		return int('z' - 'a')
	}
	return int(field[i][j] - 'a')
}

func findShortestPath(field [][]byte, start, end Point2) int {
	q := make([]Point2, 0, 1)
	visited := make(map[Point2]int)

	q = append(q, start)

	var head Point2
	for len(q) > 0 {
		head, q = q[0], q[1:]
		if head == end {
			return visited[head]
		}
		i, j := head.x, head.y
		for _, step := range STEPS4 {
			ni, nj := i+step[0], j+step[1]
			if ni < 0 || nj < 0 || ni >= len(field) || nj >= len(field[0]) {
				continue
			}
			np := Point2{ni, nj}
			if _, ok := visited[np]; ok {
				continue
			}
			if getHeight(field, ni, nj)-getHeight(field, i, j) <= 1 {
				visited[np] = visited[head] + 1
				q = append(q, np)
			}
		}
	}

	return ALOT
}
