package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay1(t *testing.T) {
	const vec string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	t.Run("1", puzzleexpect.SolveTester(Day1Part1{}, vec, "24000"))
	t.Run("2", puzzleexpect.SolveTester(Day1Part2{}, vec, "45000"))
}
