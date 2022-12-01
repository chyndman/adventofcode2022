package aoc2022

import (
	"bufio"
	"io"
	"strconv"
)

type Day1Part1 struct{}

func (_ Day1Part1) Solve(input io.Reader) (answer string, err error) {
	var acc, maxacc int
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		if 0 == len(sc.Text()) {
			if acc > maxacc {
				maxacc = acc
			}
			acc = 0
		} else {
			var n int
			n, err = strconv.Atoi(sc.Text())
			if err != nil {
				return
			}
			acc += n
		}
	}
	answer = strconv.Itoa(maxacc)
	return
}
