package demo

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

type Sum struct{}

func (_ Sum) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		var n int
		n, err = strconv.Atoi(line)
		if err != nil {
			return
		}
		acc += n
	}

	puzzle.ForEachLine(input, lineIter)

	answer = strconv.Itoa(acc)
	return
}
