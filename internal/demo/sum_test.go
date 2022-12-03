package demo

import (
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	r := strings.NewReader("1\n2\n3\n")
	answer, err := Sum{}.Solve(r)
	if nil != err || "6" != answer {
		t.Errorf("Sum Solve = %s, %v; want 6, nil", answer, err)
	}
}
