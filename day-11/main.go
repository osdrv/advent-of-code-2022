package main

import (
	"fmt"
	"os"
	"sort"
)

type Monkey struct {
	id        int
	items     []int
	op        [3]string
	test      int
	testTrue  int
	testFalse int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %+v", m.id, m.items)
}

func parseMonkey(lines []string, ix int) (*Monkey, int) {
	monkey := &Monkey{}
	ptr := 0
	_, ptr = readStr(lines[ix], 0, "Monkey ")
	monkey.id, _ = readInt(lines[ix], ptr)
	ix++
	_, ptr = readStr(lines[ix], 0, "  Starting items: ")
	monkey.items = parseInts(lines[ix][ptr:])
	ix++
	_, ptr = readStr(lines[ix], 0, "  Operation: new = ")
	monkey.op[0], ptr = readWord(lines[ix], ptr)
	monkey.op[1] = lines[ix][ptr+1 : ptr+2]
	ptr += 3
	monkey.op[2], ptr = readWord(lines[ix], ptr)
	ix++
	_, ptr = readStr(lines[ix], 0, "  Test: divisible by ")
	monkey.test, _ = readInt(lines[ix], ptr)
	ix++
	_, ptr = readStr(lines[ix], 0, "    If true: throw to monkey ")
	monkey.testTrue, _ = readInt(lines[ix], ptr)
	ix++
	_, ptr = readStr(lines[ix], 0, "    If false: throw to monkey ")
	monkey.testFalse, _ = readInt(lines[ix], ptr)
	ix++

	return monkey, ix
}

func interpWorry(w int, op [3]string) int {
	var a, b int
	if op[0] == "old" {
		a = w
	} else {
		a = parseInt(op[0])
	}
	if op[2] == "old" {
		b = w
	} else {
		b = parseInt(op[2])
	}
	if op[1] == "*" {
		return a * b
	} else if op[1] == "+" {
		return a + b
	}
	panic("should not happen")
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	monkeys := make([]*Monkey, 0, 1)

	ix := 0
	var monkey *Monkey
	for ix < len(lines) {
		monkey, ix = parseMonkey(lines, ix)
		monkeys = append(monkeys, monkey)
		ix++
	}

	debugf("monkeys: %+v", monkeys)

	inspected := make(map[int]int)

	reduceWorry := false

	tops := 1
	for _, monkey := range monkeys {
		tops *= monkey.test
	}

	for round := 0; round < 10000; round++ {
		debugf("Round %d", round)
		for id, monkey := range monkeys {
			var item int
			for len(monkey.items) > 0 {
				item, monkey.items = monkey.items[0], monkey.items[1:]
				debugf("  Monkey %d inspects item with a worry level %d", id, item)
				inspected[id]++
				worry := item
				worry = interpWorry(worry, monkey.op)
				debugf("    Worry level is %d", worry)
				if reduceWorry {
					worry /= 3
					debugf("    Worry level is divided by 3: %d", worry)
				}
				var dst int
				if worry%monkey.test == 0 {
					debugf("    Worry level is divisible by %d", monkey.test)
					dst = monkey.testTrue

				} else {
					debugf("    Worry level is not divisible by %d", monkey.test)
					dst = monkey.testFalse
				}
				debugf("    Item with worry level %d is thrown to monkey %d", worry, dst)
				monkeys[dst].items = append(monkeys[dst].items, worry%tops)
			}
		}
		debugf("monkeys: %+v", monkeys)
	}
	debugf("inspected: %+v", inspected)

	ii := make([]int, 0, 1)
	for _, i := range inspected {
		ii = append(ii, i)
	}
	sort.Ints(ii)
	monkeyBusiness := uint64(ii[len(ii)-1]) * uint64(ii[len(ii)-2])
	printf("monkeyBusiness: %d", monkeyBusiness)
}
