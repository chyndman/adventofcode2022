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

func spawnDay10Part1() (cycle func(int, *d10RegFile), getAnswer func() string) {
	acc := 0
	cycle = func(t int, cpu *d10RegFile) {
		tp := t + 1
		if 20 == tp%40 {
			acc += tp * cpu.X
		}
		return
	}
	getAnswer = func() string {
		return strconv.Itoa(acc)
	}
	return
}

func spawnDay10Part2() (cycle func(int, *d10RegFile), getAnswer func() string) {
	var crt strings.Builder
	const width int = 40
	cycle = func(t int, cpu *d10RegFile) {
		col := t % width
		if 0 == col {
			crt.Grow(width + 1)
		}

		if (cpu.X-1) <= col && col <= (cpu.X+1) {
			crt.WriteByte('#')
		} else {
			crt.WriteByte('.')
		}

		if width-1 == col {
			crt.WriteByte('\n')
		}

		return
	}
	getAnswer = func() string {
		return crt.String()
	}
	return
}

type Day10 struct {
	Spawn func() (cycle func(int, *d10RegFile), getAnswer func() string)
}

var (
	Day10Part1 = Day10{spawnDay10Part1}
	Day10Part2 = Day10{spawnDay10Part2}
)

func (p Day10) Solve(input io.Reader) (answer string, err error) {
	t := 0
	cpu := d10RegFile{1}
	cycle, getAnswer := p.Spawn()

	tick := func() {
		cycle(t, &cpu)
		t++
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
	return getAnswer(), nil
}
