package main

import (
	"os"
)

const (
	_ int = iota
	ADDX
	NOOP
)

func parseCmd(s string) [3]int {
	if startsWith(s, "addx") {
		return [3]int{ADDX, parseInt(s[5:]), 2}
	} else if startsWith(s, "noop") {
		return [3]int{NOOP, 0, 1}
	}
	panic("should not happen")
}

func execCmd(regs map[byte]int, cmd [3]int) {
	switch cmd[0] {
	case ADDX:
		regs['X'] += cmd[1]
	case NOOP:
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	input := make([][3]int, 0, len(lines))
	queue := make([][3]int, 0, 1)

	for _, line := range lines {
		input = append(input, parseCmd(line))
	}

	regs := make(map[byte]int)
	regs['X'] = 1
	strength := 0
	interesting := map[int]bool{
		20: true, 60: true, 100: true, 140: true, 180: true, 220: true,
	}

	sprite := makeByteField(6, 40)

	cycle := 1
	var head [3]int
	pos := 0
	for len(input) > 0 || len(queue) > 0 {
		debugf("cycle %d, regs: %+v", cycle, regs)

		if interesting[cycle] {
			debugf("  strength at cycle %d is %d", cycle, cycle*regs['X'])
			strength += cycle * regs['X']
		}

		posX, posY := pos%40, pos/40

		if posX >= regs['X']-1 && posX <= regs['X']+1 {
			sprite[posY][posX] = 1
		} else {
			sprite[posY][posX] = 0
		}

		if len(queue) > 0 {
			head, queue = queue[0], queue[1:]
		} else {
			head, input = input[0], input[1:]
		}
		debugf("  exec cmd: %+v", head)
		head[2]--
		if head[2] == 0 {
			execCmd(regs, head)
		} else {
			queue = append(queue, head)
		}
		debugf("  by the end of %d cycle regs are: %+v", cycle, regs)

		pos++
		pos %= 40 * 6

		cycle++
	}

	printf("total strength: %d", strength)

	print(printNumFieldWithSubs(sprite, " ", map[byte]string{
		1: "#", 0: ".",
	}))
}
