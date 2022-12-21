package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func yield(mnks map[string][]string, root string) int {
	var rec func(s string) int
	rec = func(s string) int {
		if len(mnks[s]) == 1 {
			return parseInt(mnks[s][0])
		}
		m1, m2 := rec(mnks[s][0]), rec(mnks[s][2])
		switch mnks[s][1] {
		case "+":
			return m1 + m2
		case "-":
			return m1 - m2
		case "*":
			return m1 * m2
		case "/":
			return m1 / m2
		default:
			panic("should not happen")
		}
	}

	return rec(root)
}

func simplifyStr(mnks map[string][]string, root string) string {
	var rec func(string) (string, bool)

	asStr := strconv.Itoa

	rec = func(s string) (string, bool) {
		if s == "humn" {
			return "X", false
		}
		if len(mnks[s]) == 1 {
			return mnks[s][0], true
		}
		v1, ok1 := rec(mnks[s][0])
		v2, ok2 := rec(mnks[s][2])

		if ok1 && ok2 {
			op := mnks[s][1]
			m1, m2 := parseInt(v1), parseInt(v2)
			switch op {
			case "+":
				return asStr(m1 + m2), true
			case "-":
				return asStr(m1 - m2), true
			case "*":
				return asStr(m1 * m2), true
			case "/":
				return asStr(m1 / m2), true
			default:
				panic("should not happen")
			}
		}

		return fmt.Sprintf("(%s %s %s)", v1, mnks[s][1], v2), false
	}
	res, _ := rec(root)
	return res
}

// Represents a linear equation in the form: K*X + B
type equ struct {
	K, B *big.Rat
}

func plus(a, b *big.Rat) *big.Rat {
	return big.NewRat(0, 1).Add(a, b)
}

func minus(a, b *big.Rat) *big.Rat {
	return big.NewRat(0, 1).Sub(a, b)
}

func mult(a, b *big.Rat) *big.Rat {
	return big.NewRat(0, 1).Mul(a, b)
}

func div(a, b *big.Rat) *big.Rat {
	return mult(a, big.NewRat(0, 1).Inv(b))
}

func simplify(mnks map[string][]string, root string) equ {
	var rec func(string) equ
	rec = func(s string) equ {
		if s == "humn" {
			return equ{K: big.NewRat(1, 1), B: big.NewRat(0, 1)}
		}
		if len(mnks[s]) == 1 {
			return equ{K: big.NewRat(0, 1), B: big.NewRat(int64(parseInt(mnks[s][0])), 1)}
		}

		eq1, op, eq2 := rec(mnks[s][0]), mnks[s][1], rec(mnks[s][2])

		assert(big.NewRat(0, 1).Mul(eq1.K, eq2.K).Cmp(big.NewRat(0, 1)) == 0, "no quadratic equations please!")

		switch op {
		case "+":
			return equ{K: plus(eq1.K, eq2.K), B: plus(eq1.B, eq2.B)}
		case "-":
			return equ{K: minus(eq1.K, eq2.K), B: minus(eq1.B, eq2.B)}
		case "*":
			return equ{K: plus(mult(eq1.K, eq2.B), mult(eq2.K, eq1.B)), B: mult(eq1.B, eq2.B)}
		case "/":
			assert(eq2.K.Cmp(big.NewRat(0, 1)) == 0, "no idea how to solve this")
			return equ{K: div(eq1.K, eq2.B), B: div(eq1.B, eq2.B)}
		default:
			panic("should not happen")
		}
	}

	return rec("root")
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	mnks := make(map[string][]string)

	for _, line := range lines {
		chs := strings.Split(line, " ")
		mnks[chs[0][:len(chs[0])-1]] = chs[1:]
	}

	debugf("mnks: %+v", mnks)

	num := yield(mnks, "root")
	printf("yield num: %d", num)

	debugf(simplifyStr(mnks, "root"))

	mnks["root"][1] = "-"
	eq := simplify(mnks, "root")
	x := mult(big.NewRat(-1, 1), div(eq.B, eq.K))
	printf("The root is: %d", x.Num().Int64())
}
