package main

import (
	"os"
)

func computeScore(trees [][]byte, p0 Point2) int {
	dir := complex(1, 0)

	var dst [4]int
	for turn := 0; turn < 4; turn++ {
		p := p0
		for {
			p.x += int(real(dir))
			p.y += int(imag(dir))
			if p.x < 0 || p.x >= len(trees[0]) || p.y < 0 || p.y >= len(trees) {
				break
			}
			dst[turn]++
			if trees[p.y][p.x] >= trees[p0.y][p0.x] {
				break
			}
		}
		dir *= 1i
	}

	return dst[0] * dst[1] * dst[2] * dst[3]
}

func computeVisible(trees [][]byte) map[Point2]bool {
	visible := make(map[Point2]bool)

	assert(len(trees) == len(trees[0]), "This solution only works for a square matrix")

	mod := func(turn int, p Point2) Point2 {
		for i := 0; i < turn; i++ {
			p.x, p.y = len(trees)-1-p.y, p.x
		}
		return p
	}

	data := trees
	for turn := 0; turn < 4; turn++ {
		for i := 1; i < len(data)-1; i++ {
			prev := data[i][0]
			for j := 1; j < len(data[0])-1; j++ {
				if data[i][j] > prev {
					visible[mod(turn, Point2{i, j})] = true
				}
				prev = max(prev, data[i][j])
			}
		}
		data = reverseMatHor(transposeMat(data))
	}

	return visible
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	trees := make([][]byte, 0, 1)
	for _, line := range lines {
		trees = append(trees, []byte(line))
	}

	visible := computeVisible(trees)

	debugf("visible coords: %+v", visible)

	visibleCnt := 2*len(trees) + 2*(len(trees[0])-2) + len(visible)

	printf("visible cnt: %d", visibleCnt)

	maxScore := -ALOT
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[0])-1; j++ {
			scenicScore := computeScore(trees, Point2{x: j, y: i})
			maxScore = max(scenicScore, maxScore)
		}
	}

	printf("max score: %d", maxScore)
}
