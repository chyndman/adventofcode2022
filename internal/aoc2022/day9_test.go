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
	const vecp string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	t.Run("1", puzzleexpect.SolveTester(Day9Part1, vec, "13"))
	t.Run("2", puzzleexpect.SolveTester(Day9Part2, vec, "1"))
	t.Run("2p", puzzleexpect.SolveTester(Day9Part2, vecp, "36"))
}
