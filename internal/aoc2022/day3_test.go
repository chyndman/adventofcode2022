package aoc2022

import (
	"strings"
	"testing"
)

var vecDay3 string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestDay3Part1(t *testing.T) {
	r := strings.NewReader(vecDay3)
	answer, err := Day3Part1{}.Solve(r)
	expected := "157"
	if nil != err || expected != answer {
		t.Errorf("day3.1 answer = %s, %v; want %s, nil", answer, err, expected)
	}
}
