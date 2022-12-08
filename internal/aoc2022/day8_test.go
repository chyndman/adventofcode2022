package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay8(t *testing.T) {
	const vec string = `30373
25512
65332
33549
35390
`
	t.Run("1", puzzleexpect.SolveTester(Day8Part1, vec, "21"))
	t.Run("2", puzzleexpect.SolveTester(Day8Part2, vec, "8"))
}
