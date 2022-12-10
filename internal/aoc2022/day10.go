package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
	"strings"
)

type d10RegFile struct {
	X int
}

type Day10 struct{}

var (
	Day10Part1 = Day10{}
)

func (p Day10) Solve(input io.Reader) (answer string, err error) {
	acc := 0
	t := 0
	cpu := d10RegFile{1}

	tick := func() {
		t++
		if 20 == t%40 {
			acc += t * cpu.X
		}
	}

	lineIter := func(line string) {
		instr := strings.Split(line, " ")
		switch instr[0] {
		case "addx":
			tick()
			tick()
			addend, _ := strconv.Atoi(instr[1])
			cpu.X += addend
		default:
			tick()
		}
	}

	puzzle.ForEachLine(input, lineIter)
	return strconv.Itoa(acc), nil
}
