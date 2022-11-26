package demo

import (
	"bufio"
	"io"
	"strconv"
)

type Sum struct{}

func (_ Sum) Solve(input io.Reader) (answer string, err error) {
	var acc int
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		var n int
		n, err = strconv.Atoi(sc.Text())
		if err != nil {
			return
		}
		acc += n
	}
	answer = strconv.Itoa(acc)
	return
}
