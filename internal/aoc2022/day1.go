package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

type Day1Part1 struct{}

func (_ Day1Part1) Solve(input io.Reader) (answer string, err error) {
	var acc, maxacc int
	lineIter := func (line string) {
		if 0 == len(line) {
			if acc > maxacc {
				maxacc = acc
			}
			acc = 0
		} else {
			var n int
			n, err = strconv.Atoi(line)
			if err != nil {
				return
			}
			acc += n
		}
	}

	puzzle.ForEachLine(input, lineIter)
	lineIter("")

	answer = strconv.Itoa(maxacc)
	return
}

type Day1Part2 struct{}

func (_ Day1Part2) Solve(input io.Reader) (answer string, err error) {
	var acc int
	var maxaccs [3]int
	lineIter := func (line string) {
		if 0 == len(line) {
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
			n, err = strconv.Atoi(line)
			if err != nil {
				return
			}
			acc += n
		}
	}
	
	puzzle.ForEachLine(input, lineIter)
	lineIter("")

	answer = strconv.Itoa(maxaccs[0] + maxaccs[1] + maxaccs[2])
	return
}
