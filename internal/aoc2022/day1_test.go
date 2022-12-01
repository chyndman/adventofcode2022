package aoc2022

import (
	"testing"
	"strings"
)

func TestDay1Part1(t *testing.T) {
	r := strings.NewReader(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`)
	answer, err := Day1Part1{}.Solve(r)
	expected := "24000"
	if nil != err || expected != answer {
		t.Errorf("Day1Part1 answer = %s, %v; want %s, nil", answer, err, expected)
	}
}
