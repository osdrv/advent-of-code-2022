package main

import (
	"bytes"
	"os"
)

func parseMove(s string) [3]int {
	var am, from, to int
	ptr := 0
	_, ptr = readStr(s, ptr, "move ")
	am, ptr = readInt(s, ptr)
	_, ptr = readStr(s, ptr, " from ")
	from, ptr = readInt(s, ptr)
	_, ptr = readStr(s, ptr, " to ")
	to, ptr = readInt(s, ptr)
	return [3]int{am, from, to}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	stacks := make([][]byte, 0, 1)
	ix := 0
	for ix < len(lines) {
		if len(lines[ix]) == 0 {
			break
		}
		stacks = append(stacks, []byte(lines[ix]))
		ix++
	}
	ix++
	moves := make([][3]int, 0, 1)
	for ix < len(lines) {
		moves = append(moves, parseMove(lines[ix]))
		ix++
	}

	printf("stacks: %+v", stacks)
	printf("moves: %+v", moves)

	// false for part1, true for part2
	moveBlocks := true

	for _, move := range moves {
		am, from, to := move[0], move[1], move[2]
		if moveBlocks {
			stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-am:]...)
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-am]
		} else {
			for i := 0; i < am; i++ {
				stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
				stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
			}
		}

		debugf("stacks: %+v", stacks)
	}

	var b bytes.Buffer
	for _, stack := range stacks {
		b.WriteByte(stack[len(stack)-1])
	}

	printf("result: %s", b.String())
}
