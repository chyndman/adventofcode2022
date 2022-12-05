package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"regexp"
	"strconv"
)

func supplyStackPush(stacks map[int]string, idx int, crates string) {
	stk, exists := stacks[idx]
	if !exists {
		stk = ""
	}
	stacks[idx] = crates + stk
}

func supplyStackPop(stacks map[int]string, idx, count int, rev bool) (crates string) {
	crates = ""
	stk, exists := stacks[idx]
	if exists {
		for i := 0; i < count; i++ {
			if rev {
				crates = string(stk[i]) + crates
			} else {
				crates = crates + string(stk[i])
			}
		}
		stacks[idx] = stk[count:]
	}
	return
}

var supplyStackProcStmtPat = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func supplyStackProcStmtDecode(line string) (count, stkFrom, stkTo int) {
	submatch := supplyStackProcStmtPat.FindStringSubmatch(line)
	count, _ = strconv.Atoi(submatch[1])
	stkFrom, _ = strconv.Atoi(submatch[2])
	stkTo, _ = strconv.Atoi(submatch[3])
	return
}

func solveDay5Common(input io.Reader, rev bool) (answer string, err error) {
	stacks := make(map[int]string)
	maxStkIdx := 0
	state := "stacks"
	lineIter := func(line string) {
		switch state {
		case "stacks":
			if '1' == line[1] {
				state = "stacks end"
			} else {
				stkIdx := 0
				for i := 1; i < len(line); i += 4 {
					stkIdx++
					c := line[i]
					if 'A' <= c && c <= 'Z' {
						stacks[stkIdx] = stacks[stkIdx] + string(c)
					}
				}
				maxStkIdx = stkIdx
			}
		case "stacks end":
			state = "procedure"
		case "procedure":
			count, stkFrom, stkTo := supplyStackProcStmtDecode(line)
			supplyStackPush(stacks, stkTo, supplyStackPop(stacks, stkFrom, count, rev))
		}
	}

	puzzle.ForEachLine(input, lineIter)

	for i := 1; i <= maxStkIdx; i++ {
		stk := stacks[i]
		if 0 < len(stk) {
			answer += stk[0:1]
		}
	}

	return
}

type Day5Part1 struct{}

func (_ Day5Part1) Solve(input io.Reader) (answer string, err error) {
	return solveDay5Common(input, true)
}

type Day5Part2 struct{}

func (_ Day5Part2) Solve(input io.Reader) (answer string, err error) {
	return solveDay5Common(input, false)
}
