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
	// last elf has no blank line
	if acc > maxacc {
		maxacc = acc
	}
	answer = strconv.Itoa(maxacc)
	return
}

type Day1Part2 struct{}

func (_ Day1Part2) Solve(input io.Reader) (answer string, err error) {
	var acc int
	var maxaccs [3]int
	sc := bufio.NewScanner(input)
	lineIter := func (txt string) {
		if 0 == len(txt) {
			for i := 0; i < len(maxaccs); i++ {
				if acc > maxaccs[i] {
					for j := len(maxaccs) - 1; j > i; j-- {
						maxaccs[j] = maxaccs[j - 1]
					}
					maxaccs[i] = acc
					break
				}
			}
			acc = 0
		} else {
			var n int
			n, err = strconv.Atoi(txt)
			if err != nil {
				return
			}
			acc += n
		}
	}
	
	for sc.Scan() {
		lineIter(sc.Text())
	}
	lineIter("")

	answer = strconv.Itoa(maxaccs[0] + maxaccs[1] + maxaccs[2])
	return
}
