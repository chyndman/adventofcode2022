package puzzle

import (
	"io"
	"bufio"
)

type Puzzle interface {
	Solve(input io.Reader) (answer string, err error)
}

func ForEachLine(input io.Reader, fn func (line string)) {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		fn(sc.Text())
	}
}
