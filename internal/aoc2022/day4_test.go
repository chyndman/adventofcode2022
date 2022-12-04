package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay4(t *testing.T) {
	const vec string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`
	t.Run("1", puzzleexpect.SolveTester(Day4Part1{}, vec, "2"))
}
