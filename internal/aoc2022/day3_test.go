package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay3(t *testing.T) {
	const vec string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	t.Run("1", puzzleexpect.SolveTester(Day3Part1{}, vec, "157"))
	t.Run("2", puzzleexpect.SolveTester(Day3Part2{}, vec, "70"))
}
