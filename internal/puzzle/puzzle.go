package puzzle

import (
	"bufio"
	"io"
	"regexp"
)

type Puzzle interface {
	Solve(input io.Reader) (answer string, err error)
}

func ForEachLine(input io.Reader, fn func(line string)) {
	for sc := bufio.NewScanner(input); sc.Scan(); {
		fn(sc.Text())
	}
}

func ForEachLineSubmatches(input io.Reader, pat string, fn func(submatches []string)) {
	for r, sc := regexp.MustCompile(pat), bufio.NewScanner(input); sc.Scan(); {
		fn(r.FindStringSubmatch(sc.Text()))
	}
}
