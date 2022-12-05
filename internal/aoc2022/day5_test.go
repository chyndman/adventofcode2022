package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay5(t *testing.T) {
	const vec string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	t.Run("1", puzzleexpect.SolveTester(Day5Part1{}, vec, "CMZ"))
}
