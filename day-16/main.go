package main

import (
	"os"
)

type Valve struct {
	n string
	r int
	d []string
}

const (
	MAXT = 30
)

func findMaxRate(vv []*Valve) int {
	vix := make(map[string]int)
	for ix, v := range vv {
		vix[v.n] = ix
	}

	memo := make(map[[3]int]int)

	open := make(map[int]bool)

	var visit func(int, int, int) int
	visit = func(p, time, rate int) int {
		if time <= 0 {
			return 0
		}
		key := [3]int{p, time, rate}
		if prev, ok := memo[key]; ok {
			return prev
		}
		mr := time * rate

		if !open[p] && vv[p].r > 0 {
			open[p] = true
			mr = max(mr, rate+visit(p, time-1, rate+vv[p].r))
			delete(open, p)
		}

		for _, next := range vv[p].d {
			mr = max(mr, rate+visit(vix[next], time-1, rate))

		}
		memo[key] = mr
		return mr
	}

	return visit(vix["AA"], MAXT, 0)
}

func findMaxRate2(vv []*Valve) int {
	vix := make(map[string]int)
	for ix, v := range vv {
		vix[v.n] = ix
	}

	var open uint64

	memo := make(map[[4]int]int)

	var visit func(int, int, int, int) int
	visit = func(a, b int, time int, rate int) int {
		if time <= 0 {
			return 0
		}

		key := [4]int{a, b, time, rate}
		if prev, ok := memo[key]; ok {
			return prev
		}

		if a > b {
			a, b = b, a
		}
		mr := time * rate
		if (open&(1<<a) == 0) && (open&(1<<b) == 0) && a != b {
			if vv[a].r > 0 && vv[b].r > 0 {
				open |= 1 << a
				open |= 1 << b
				mr = max(mr, rate+visit(a, b, time-1, rate+vv[a].r+vv[b].r))
				open &= ^(1 << a)
				open &= ^(1 << b)
			}
		}
		if (open&(1<<a) == 0) && vv[a].r > 0 {
			open |= 1 << a
			for _, bnext := range vv[b].d {
				mr = max(mr, rate+visit(a, vix[bnext], time-1, rate+vv[a].r))
			}
			open &= ^(1 << a)
		}
		if a != b {
			if (open&(1<<b) == 0) && vv[b].r > 0 {
				open |= 1 << b
				for _, anext := range vv[a].d {
					mr = max(mr, rate+visit(vix[anext], b, time-1, rate+vv[b].r))
				}
				open &= ^(1 << b)
			}
		}
		for _, anext := range vv[a].d {
			for _, bnext := range vv[b].d {
				mr = max(mr, rate+visit(vix[anext], vix[bnext], time-1, rate))
			}
		}

		memo[key] = mr

		return mr
	}

	return visit(vix["AA"], vix["AA"], MAXT-4, 0)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	vv := make([]*Valve, 0, 1)
	for _, line := range lines {
		var name string
		var rate int
		var dummy1, dummy2, dummy3 string
		var dirs []string

		Scanf(line, "Valve {string} has flow rate={int}; {string} {string} to {string} {[]string}", &name, &rate, &dummy1, &dummy3, &dummy2, &dirs)
		vv = append(vv, &Valve{n: name, r: rate, d: dirs})
		debugf("valve: %+v", vv[len(vv)-1])
	}

	maxR := findMaxRate(vv)
	printf("part 1: %d", maxR)

	maxR2 := findMaxRate2(vv)
	printf("part 2: %d", maxR2)
}
