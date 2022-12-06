package aoc2022

import (
	"bufio"
	"io"
	"strconv"
)

func d6LetterIdx(c byte) int {
	if 'a' <= c && c <= 'z' {
		return int(c - 'a')
	} else {
		return -1
	}
}

type Day6 struct {
	windowSize int
}

var (
	Day6Part1 = Day6{4}
	Day6Part2 = Day6{14}
)

func (p Day6) Solve(input io.Reader) (answer string, err error) {
	unique, ring, hist := 0, make([]byte, p.windowSize), make([]int, 26)
	for i, r := 0, bufio.NewReader(input); nil == err && "" == answer; i++ {
		ringIdx := i % p.windowSize

		var newLetter byte
		newLetter, err = r.ReadByte()
		newIdx := d6LetterIdx(newLetter)
		if newIdx < 0 {
			break
		}

		if i >= p.windowSize {
			oldLetter := ring[ringIdx]
			oldIdx := d6LetterIdx(oldLetter)
			hist[oldIdx]--
			if 0 == hist[oldIdx] {
				unique--
			}
		}

		ring[ringIdx] = newLetter
		hist[newIdx]++
		if 1 == hist[newIdx] {
			unique++
			if p.windowSize == unique {
				answer = strconv.Itoa(i + 1)
			}
		}
	}
	return
}
