package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay12(t *testing.T) {
	const vec string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

	t.Run("1", puzzleexpect.SolveTester(Day12Part1, vec, "31"))
	t.Run("2", puzzleexpect.SolveTester(Day12Part2, vec, "29"))
}
