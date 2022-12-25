package main

import (
	"os"
)

func snafuToDec(s string) int64 {
	mult := int64(1)
	num := int64(0)
	for ptr := len(s) - 1; ptr >= 0; ptr-- {
		var v int64
		switch s[ptr] {
		case '2':
			v = 2
		case '1':
			v = 1
		case '0':
			v = 0
		case '-':
			v = -1
		case '=':
			v = -2
		default:
			panic("oopsie")
		}
		num += v * mult
		mult *= 5
	}
	return num
}

func decToSnafu(num int64) string {
	if num == 0 {
		return "0"
	}
	b := make([]byte, 0, 1)
	for num > 0 {
		rem := num % 5
		switch rem {
		case 0:
			b = append(b, '0')
		case 1:
			b = append(b, '1')
		case 2:
			b = append(b, '2')
		case 3:
			b = append(b, '=')
			num += 2
		case 4:
			b = append(b, '-')
			num += 1
		}
		num /= 5
	}
	for i := 0; i <= (len(b)-1)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return string(b)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	sum := int64(0)
	for _, line := range lines {
		num := snafuToDec(line)
		dec := decToSnafu(num)
		debugf("%s -> %d -> %s", line, num, dec)
		sum += num
	}

	printf("sum: %d", sum)

	printf("sum -> snafu: %s", decToSnafu(sum))
}
