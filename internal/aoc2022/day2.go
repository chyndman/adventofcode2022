package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"regexp"
	"strconv"
)

var rpsPat = regexp.MustCompile("([A-C]) ([X-Z])")

const (
	rock     = 0
	paper    = 1
	scissors = 2

	loss = -1
	draw = 0
	win  = 1
)

func rpsResult(x, y int) (r int) {
	r = draw
	xp := (x + 1) % 3
	yp := (y + 1) % 3
	if xp == y {
		r = win
	} else if yp == x {
		r = loss
	}
	return
}

func rpsDecode(letter string) (num int) {
	switch letter {
	case "A", "X":
		num = rock
	case "B", "Y":
		num = paper
	case "C", "Z":
		num = scissors
	}
	return
}

func rpsResultDecode(letter string) (num int) {
	switch letter {
	case "X":
		num = loss
	case "Y":
		num = draw
	case "Z":
		num = win
	}
	return
}

func rpsHandScore(h int) int {
	return h + 1
}

func rpsResultScore(r int) (sc int) {
	return (r + 1) * 3
}

type Day2Part1 struct{}

func (_ Day2Part1) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		submatch := rpsPat.FindStringSubmatch(line)
		opp, me := rpsDecode(submatch[1]), rpsDecode(submatch[2])
		meScore := rpsHandScore(me)
		result := rpsResult(opp, me)
		resultScore := rpsResultScore(result)
		acc += (meScore + resultScore)
	}

	puzzle.ForEachLine(input, lineIter)

	answer = strconv.Itoa(acc)
	return
}

type Day2Part2 struct{}

func (_ Day2Part2) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		submatch := rpsPat.FindStringSubmatch(line)
		opp, result := rpsDecode(submatch[1]), rpsResultDecode(submatch[2])
		var me int
		for i := rock; i <= scissors; i++ {
			if rpsResult(opp, i) == result {
				me = i
				break
			}
		}
		meScore := rpsHandScore(me)
		resultScore := rpsResultScore(result)
		acc += (meScore + resultScore)
	}

	puzzle.ForEachLine(input, lineIter)

	answer = strconv.Itoa(acc)
	return
}
