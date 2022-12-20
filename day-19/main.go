package main

import (
	"os"
)

type cost struct {
	ore, clay, obsidian, geode int
}

func parseBP(s string) [4]cost {
	var n1, n2, n3, n4, n5, n6, n7 int
	Scanf(s, "Blueprint {int}: Each ore robot costs {int} ore. Each clay robot costs {int} ore. Each obsidian robot costs {int} ore and {int} clay. Each geode robot costs {int} ore and {int} obsidian.", &n1, &n2, &n3, &n4, &n5, &n6, &n7)
	var bp [4]cost
	bp[ORE] = cost{ore: n2}
	bp[CLAY] = cost{ore: n3}
	bp[OBSIDIAN] = cost{ore: n4, clay: n5}
	bp[GEODE] = cost{ore: n6, obsidian: n7}
	return bp
}

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
)

func bestQuality(bp [4]cost, time int) int {
	maxOre := max(bp[ORE].ore, max(bp[CLAY].ore, max(bp[OBSIDIAN].ore, bp[GEODE].ore)))

	memo := make(map[[8]int]int)
	var rec func(t int, robots [4]int, budget cost) int
	rec = func(t int, robots [4]int, budget cost) int {
		if t == 0 {
			return budget.geode
		}

		key := [8]int{t, robots[0], robots[1], robots[2], robots[3], budget.clay, budget.obsidian, budget.ore}
		if prev, ok := memo[key]; ok {
			return prev
		}

		// do nothing

		mq := 0

		canMake := [4]bool{
			budget.ore >= bp[ORE].ore,  // ORE
			budget.ore >= bp[CLAY].ore, // CLAY
			budget.ore >= bp[OBSIDIAN].ore && budget.clay >= bp[OBSIDIAN].clay,   // OBSIDIAN
			budget.ore >= bp[GEODE].ore && budget.obsidian >= bp[GEODE].obsidian, // GEODE
		}

		if canMake[GEODE] {
			mq = max(mq, rec(t-1, [4]int{robots[ORE], robots[CLAY], robots[OBSIDIAN], robots[GEODE] + 1}, cost{
				ore:      budget.ore + robots[ORE] - bp[GEODE].ore,
				clay:     budget.clay + robots[CLAY] - bp[GEODE].clay,
				obsidian: budget.obsidian + robots[OBSIDIAN] - bp[GEODE].obsidian,
				geode:    budget.geode + robots[GEODE],
			}))
		}

		if canMake[ORE] && robots[ORE] < maxOre {
			mq = max(mq, rec(t-1, [4]int{robots[ORE] + 1, robots[CLAY], robots[OBSIDIAN], robots[GEODE]}, cost{
				ore:      budget.ore + robots[ORE] - bp[ORE].ore,
				clay:     budget.clay + robots[CLAY] - bp[ORE].clay,
				obsidian: budget.obsidian + robots[OBSIDIAN] - bp[ORE].obsidian,
				geode:    budget.geode + robots[GEODE],
			}))
		}

		if canMake[CLAY] && robots[CLAY] < bp[OBSIDIAN].clay {
			mq = max(mq, rec(t-1, [4]int{robots[ORE], robots[CLAY] + 1, robots[OBSIDIAN], robots[GEODE]}, cost{
				ore:      budget.ore + robots[ORE] - bp[CLAY].ore,
				clay:     budget.clay + robots[CLAY] - bp[CLAY].clay,
				obsidian: budget.obsidian + robots[OBSIDIAN] - bp[CLAY].obsidian,
				geode:    budget.geode + robots[GEODE],
			}))
		}

		if canMake[OBSIDIAN] && robots[OBSIDIAN] < bp[GEODE].obsidian {
			mq = max(mq, rec(t-1, [4]int{robots[ORE], robots[CLAY], robots[OBSIDIAN] + 1, robots[GEODE]}, cost{
				ore:      budget.ore + robots[ORE] - bp[OBSIDIAN].ore,
				clay:     budget.clay + robots[CLAY] - bp[OBSIDIAN].clay,
				obsidian: budget.obsidian + robots[OBSIDIAN] - bp[OBSIDIAN].obsidian,
				geode:    budget.geode + robots[GEODE],
			}))
		}

		if !(canMake[ORE] && canMake[CLAY] && canMake[OBSIDIAN] && canMake[GEODE]) {
			mq = max(mq, rec(t-1, robots, cost{
				ore:      budget.ore + robots[ORE],
				clay:     budget.clay + robots[CLAY],
				obsidian: budget.obsidian + robots[OBSIDIAN],
				geode:    budget.geode + robots[GEODE],
			}))
		}

		memo[key] = mq

		return memo[key]
	}

	return rec(time, [4]int{1, 0, 0, 0}, cost{})
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	bps := make([][4]cost, 0, len(lines))

	for _, line := range lines {
		bp := parseBP(line)
		debugf("blueprint: %+v", bp)
		bps = append(bps, bp)
	}

	printf("%+v", bps)

	sum := 0
	for ix, bp := range bps {
		q := bestQuality(bp, 24)
		printf("bp %d bq %d", ix+1, q)
		sum += (ix + 1) * q
	}

	printf("sum: %d", sum)

	mg := 1

	for i := 0; i < min(len(bps), 3); i++ {
		bq := bestQuality(bps[i], 32)
		printf("id: %d, bq: %d", i+1, bq)
		mg *= bq
	}

	printf("mg: %d", mg)
}
