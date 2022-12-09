package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"image"
	"io"
	"strconv"
)

type Day9 struct {
	KnotCount int
}

var (
	Day9Part1 = Day9{2}
	Day9Part2 = Day9{10}
)

func (p Day9) Solve(input io.Reader) (answer string, err error) {
	acc := 1
	knots := make([]image.Point, p.KnotCount)
	tvisits := make(map[image.Point]bool)
	tvisits[image.Point{}] = true

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
			knots[0] = knots[0].Add(mv)

			for i := 1; i < len(knots); i++ {
				diff := knots[i-1].Sub(knots[i])

				mv := false
				tmv := diff
				if 2 == diff.X {
					tmv.X = 1
					mv = true
				}
				if -2 == diff.X {
					tmv.X = -1
					mv = true
				}
				if 2 == diff.Y {
					tmv.Y = 1
					mv = true
				}
				if -2 == diff.Y {
					tmv.Y = -1
					mv = true
				}

				if mv {
					knots[i] = knots[i].Add(tmv)
					if i == len(knots)-1 {
						if _, exists := tvisits[knots[i]]; !exists {
							acc++
							tvisits[knots[i]] = true
						}
					}
				}
			}
		}
	}

	puzzle.ForEachLineSubmatches(input, `(L|R|D|U) (\d+)`, lsmIter)
	return strconv.Itoa(acc), nil
}
