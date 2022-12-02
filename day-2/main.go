package main

import (
	"os"
)

func play(a, b int) int {
	/*
		Rc  Pp  Sc
		a = 0
		b = 0 -> 1  a+b -> 0  b-(a+1)%3 -> -1  (a+1)%3 -> 1  (b+1)%3 -> 1
		b = 1 -> 2  a+b -> 1  b-(a+1)%3 -> 0   (a+1)%3 -> 1  (b+1)%3 -> 2
		b = 2 -> 0  a+b -> 2  b-(a+1)%3 -> 1   (a+1)%3 -> 1  (b+1)%3 -> 0

		a = 1
		b = 0 -> 0  a+b -> 1  b-(a+1)%3 -> -2  (a+1)%3 -> 2  (b+1)%3 -> 1
		b = 1 -> 1  a+b -> 2  b-(a+1)%3 -> -1  (a+1)%3 -> 2  (b+1)%3 -> 2
		b = 2 -> 2  a+b -> 3  b-(a+1)%3 -> 0   (a+1)%3 -> 2  (b+1)%3 -> 0

		a = 2
		b = 0 -> 2  a+b -> 2  b-(a+1)%3 -> 0   (a+1)%3 -> 0  (b+1)%3 -> 1
		b = 1 -> 0  a+b -> 3  b-(a+1)%3 -> 1   (a+1)%3 -> 0  (b+1)%3 -> 2
		b = 2 -> 1  a+b -> 4  b-(a+1)%3 -> 2   (a+1)%3 -> 0  (b+1)%3 -> 0
	*/
	return 3*(((b-(a+1)%3)+2)%3) + (b + 1)
}

func play2(a, b int) int {
	return 3*b + ((a+b-1+3)%3 + 1)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	score := 0
	score2 := 0
	for _, line := range lines {
		a, b := int(line[0]-'A'), int(line[2]-'X')
		score += play(a, b)
		score2 += play2(a, b)
	}
	printf("score: %d", score)
	printf("score2: %d", score2)
}
