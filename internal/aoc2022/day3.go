package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

type Day3Part1 struct{}

func ruckLetter2Prio(ltr byte) (prio int) {
	switch {
	case 'a' <= ltr && ltr <= 'z':
		prio = 1 + int(ltr-'a')
	case 'A' <= ltr && ltr <= 'Z':
		prio = 27 + int(ltr-'A')
	}
	return
}

var (
	ruckPrioMin = ruckLetter2Prio('a')
	ruckPrioMax = ruckLetter2Prio('Z')
)

func ruckItemMask(compart string) (mask uint64) {
	for i := 0; i < len(compart); i++ {
		mask |= 1 << ruckLetter2Prio(compart[i])
	}
	return
}

func (_ Day3Part1) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		halflen := len(line) / 2
		maskl, maskr := ruckItemMask(line[:halflen]), ruckItemMask(line[halflen:])
		masklr := maskl & maskr
		for p := ruckPrioMin; p <= ruckPrioMax; p++ {
			if 0 != masklr&(1<<p) {
				acc += p
				break
			}
		}
	}

	puzzle.ForEachLine(input, lineIter)
	answer = strconv.Itoa(acc)
	return
}
