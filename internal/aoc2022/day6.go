package aoc2022

import (
	"bufio"
	"io"
	"strconv"
)

func d6LetterIdx(c byte) int {
	if 'a' <= c && c <= 'z' {
		return int(c-'a')
	} else {
		return -1
	}
}

func solveDay6Common(input io.Reader, windowSize int) (answer string, err error) {
	unique, ring, hist := 0, make([]byte, windowSize), make([]int, 26)
	for i, r := 0, bufio.NewReader(input); nil == err && "" == answer; i++ {
		ringIdx := i % windowSize

		var newLetter byte
		newLetter, err = r.ReadByte()
		newIdx := d6LetterIdx(newLetter)
		if newIdx < 0 {
			break
		}

		if i >= windowSize {
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
			if windowSize == unique {
				answer = strconv.Itoa(i + 1)
			}
		}
	}
	return
}

type Day6Part1 struct{}

func (_ Day6Part1) Solve(input io.Reader) (string, error) {
	return solveDay6Common(input, 4)
}

type Day6Part2 struct{}

func (_ Day6Part2) Solve(input io.Reader) (string, error) {
	return solveDay6Common(input, 14)
}
