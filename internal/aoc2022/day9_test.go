package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay9(t *testing.T) {
	const vec string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	t.Run("1", puzzleexpect.SolveTester(Day9Part1, vec, "13"))
}
