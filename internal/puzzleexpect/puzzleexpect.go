package puzzleexpect

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"strings"
	"testing"
)

func SolveTester(pz puzzle.Puzzle, input, expected string) func(t *testing.T) {
	return func(t *testing.T) {
		r := strings.NewReader(input)
		answer, err := pz.Solve(r)
		if nil != err || expected != answer {
			t.Errorf("answer = %s, %v; want %s, <nil>", answer, err, expected)
		}
	}
}
