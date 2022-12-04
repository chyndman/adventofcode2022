package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"regexp"
	"strconv"
)

var campSectionRangePairPat = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func campSectionRangePairDecode(line string) (xa, xb, ya, yb int) {
	submatch := campSectionRangePairPat.FindStringSubmatch(line)
	xa, _ = strconv.Atoi(submatch[1])
	xb, _ = strconv.Atoi(submatch[2])
	ya, _ = strconv.Atoi(submatch[3])
	yb, _ = strconv.Atoi(submatch[4])
	return
}

type Day4Part1 struct{}

func (_ Day4Part1) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		xa, xb, ya, yb := campSectionRangePairDecode(line)
		mina, maxb := xa, xb
		if ya < xa {
			mina = ya
		}
		if yb > xb {
			maxb = yb
		}

		if (xa == mina && xb == maxb) || (ya == mina && yb == maxb) {
			acc++
		}
	}

	puzzle.ForEachLine(input, lineIter)
	answer = strconv.Itoa(acc)
	return
}

type Day4Part2 struct{}

func (_ Day4Part2) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		xa, xb, ya, yb := campSectionRangePairDecode(line)
		mina, maxb := xa, xb
		if ya < xa {
			mina = ya
		}
		if yb > xb {
			maxb = yb
		}

		if maxb-mina <= (xb-xa)+(yb-ya) {
			acc++
		}
	}

	puzzle.ForEachLine(input, lineIter)
	answer = strconv.Itoa(acc)
	return
}
