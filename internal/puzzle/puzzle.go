package puzzle

import (
	"io"
)

type Puzzle interface {
	Solve(input io.Reader) (answer string, err error)
}
