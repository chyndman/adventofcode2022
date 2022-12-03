package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay2(t *testing.T) {
	const vec string = `A Y
B X
C Z
`
	t.Run("1", puzzleexpect.SolveTester(Day2Part1{}, vec, "15"))
	t.Run("2", puzzleexpect.SolveTester(Day2Part2{}, vec, "12"))
}
