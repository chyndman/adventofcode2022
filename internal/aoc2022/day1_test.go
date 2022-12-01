package aoc2022

import (
	"testing"
	"strings"
)

var vec string = `1000
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

func TestDay1Part1(t *testing.T) {
	r := strings.NewReader(vec)
	answer, err := Day1Part1{}.Solve(r)
	expected := "24000"
	if nil != err || expected != answer {
		t.Errorf("Day1Part1 answer = %s, %v; want %s, nil", answer, err, expected)
	}
}

func TestDay1Part2(t *testing.T) {
	r := strings.NewReader(vec)
	answer, err := Day1Part2{}.Solve(r)
	expected := "45000"
	if nil != err || expected != answer {
		t.Errorf("Day1Part2 answer = %s, %v; want %s, nil", answer, err, expected)
	}
}
