package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay6(t *testing.T) {
	const vec string = `mjqjpqmgbljsphdztnvjfqwrcgsmlb
`
	t.Run("1", puzzleexpect.SolveTester(Day6Part1{}, vec, "7"))
	t.Run("2", puzzleexpect.SolveTester(Day6Part2{}, vec, "19"))
}
