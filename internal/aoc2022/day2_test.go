package aoc2022

import (
	"testing"
	"strings"
)

var vecDay2 string = `A Y
B X
C Z
`

func TestDay2Part1(t *testing.T) {
	r := strings.NewReader(vecDay2)
	answer, err := Day2Part1{}.Solve(r)
	expected := "15"
	if nil != err || expected != answer {
		t.Errorf("Day2Part1 answer = %s, %v; want %s, nil", answer, err, expected)
	}
}
