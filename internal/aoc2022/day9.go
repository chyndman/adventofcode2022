package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"image"
	"io"
	"strconv"
)

type Day9 struct{}

var (
	Day9Part1 = Day9{}
)

func (p Day9) Solve(input io.Reader) (answer string, err error) {
	acc := 1
	var head, tail image.Point
	tvisits := make(map[image.Point]bool)
	tvisits[tail] = true

	lsmIter := func(submatches []string) {
		mv := image.Point{}

		switch submatches[1][0] {
		case 'L':
			mv.X--
		case 'R':
			mv.X++
		case 'D':
			mv.Y--
		case 'U':
			mv.Y++
		}

		for dist, _ := strconv.Atoi(submatches[2]); dist > 0; dist-- {
			head = head.Add(mv)
			diff := head.Sub(tail)

			mv := true
			tmv := diff
			switch {
			case 2 == diff.X:
				tmv.X = 1
			case -2 == diff.X:
				tmv.X = -1
			case 2 == diff.Y:
				tmv.Y = 1
			case -2 == diff.Y:
				tmv.Y = -1
			default:
				mv = false
			}

			if mv {
				tail = tail.Add(tmv)
				if _, exists := tvisits[tail]; !exists {
					acc++
					tvisits[tail] = true
				}
			}
		}
	}

	puzzle.ForEachLineSubmatches(input, `(L|R|D|U) (\d+)`, lsmIter)
	return strconv.Itoa(acc), nil
}
