package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay14(t *testing.T) {
	const vec string = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

	t.Run("1", puzzleexpect.SolveTester(Day14Part1{}, vec, "24"))
}
