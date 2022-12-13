package main

import (
	"os"
	"sort"
)

func parseList(s string, ptr int) ([]any, int) {
	list := make([]any, 0, 1)
	ptr = consume(s, ptr, '[')
	var el any
	for ptr < len(s) {
		if match(s, ptr, ']') {
			break
		}
		if match(s, ptr, '[') {
			el, ptr = parseList(s, ptr)
		} else {
			el, ptr = readInt(s, ptr)
		}
		list = append(list, el)
		if match(s, ptr, ',') {
			ptr = consume(s, ptr, ',')
		}
	}
	ptr = consume(s, ptr, ']')
	return list, ptr
}

func compareLists(left, right any) int {
	intleft, lok := left.(int)
	intright, rok := right.(int)

	if lok && rok {
		return intleft - intright
	}

	if lok && !rok {
		return compareLists([]any{left}, right)
	} else if !lok && rok {
		return compareLists(left, []any{right})
	}
	aleft := left.([]any)
	aright := right.([]any)
	lix, rix := 0, 0
	for lix < len(aleft) && rix < len(aright) {
		if cmp := compareLists(aleft[lix], aright[rix]); cmp != 0 {
			return cmp
		}
		lix++
		rix++
	}
	return (len(aleft) - lix) - (len(aright) - rix)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	packets := make([][]any, 0, 1)

	DIV_2 := []any{[]any{2}}
	DIV_6 := []any{[]any{6}}

	pairIx := 1
	sumIx := 0
	for ix := 0; ix < len(lines); ix += 3 {
		left, _ := parseList(lines[ix], 0)
		right, _ := parseList(lines[ix+1], 0)

		debugf("left: %+v", left)
		debugf("right: %+v", right)

		cmp := compareLists(left, right)
		debugf("cmp: %d", cmp)

		if cmp < 0 {
			sumIx += pairIx
		}
		pairIx++

		packets = append(packets, left, right)
	}

	packets = append(packets, DIV_2, DIV_6)

	printf("sum ix: %d", sumIx)

	sort.Slice(packets, func(a, b int) bool {
		return compareLists(packets[a], packets[b]) < 0
	})

	debugf("packets: %+v", packets)

	var ix2, ix6 int
	for i := 0; i < len(packets); i++ {
		if compareLists(DIV_2, packets[i]) == 0 {
			ix2 = i + 1
		} else if compareLists(DIV_6, packets[i]) == 0 {
			ix6 = i + 1
		}
	}

	printf("the key is: %d", ix2*ix6)
}
