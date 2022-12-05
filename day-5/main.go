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

func parseStacks(lines []string) [][]byte {
	fld := make([][]byte, 0, len(lines))
	for _, line := range lines {
		fld = append(fld, []byte(line))
	}
	fld = transposeMat(reverseMatVer(fld))

	var b bytes.Buffer
	for _, ff := range fld {
		b.WriteString(string(ff))
		b.WriteByte('\n')
	}

	println(b.String())

	stacks := make([][]byte, 0, 1)
	for ix := 1; ix < len(fld); ix += 4 {
		ss := make([]byte, 0, 1)
		jx := 1
		for jx < len(fld[ix]) && fld[ix][jx] != ' ' {
			ss = append(ss, fld[ix][jx])
			jx++
		}
		stacks = append(stacks, ss)
	}
	return stacks
}

func main() {
	f, err := os.Open("INPUT.0")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	ix := 0
	for ix < len(lines) && len(lines[ix]) != 0 {
		ix++
	}
	stacks := parseStacks(lines[:ix])

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
