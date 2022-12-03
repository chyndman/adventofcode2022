package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

func ruckLetter2Prio(ltr byte) (prio int) {
	switch {
	case 'a' <= ltr && ltr <= 'z':
		prio = 1 + int(ltr-'a')
	case 'A' <= ltr && ltr <= 'Z':
		prio = 27 + int(ltr-'A')
	}
	return
}

var (
	ruckPrioMin = ruckLetter2Prio('a')
	ruckPrioMax = ruckLetter2Prio('Z')
)

func ruckItemMask(compart string) (mask uint64) {
	for i := 0; i < len(compart); i++ {
		mask |= 1 << ruckLetter2Prio(compart[i])
	}
	return
}

func ruckPrioScan(mask uint64) (prio int) {
	for p := ruckPrioMin; p <= ruckPrioMax; p++ {
		if 0 != mask&(1<<p) {
			prio = p
			break
		}
	}
	return
}

type Day3Part1 struct{}

func (_ Day3Part1) Solve(input io.Reader) (answer string, err error) {
	var acc int
	lineIter := func(line string) {
		halflen := len(line) / 2
		maskl, maskr := ruckItemMask(line[:halflen]), ruckItemMask(line[halflen:])
		masklr := maskl & maskr
		acc += ruckPrioScan(masklr)
	}

	puzzle.ForEachLine(input, lineIter)
	answer = strconv.Itoa(acc)
	return
}

type Day3Part2 struct{}

func (_ Day3Part2) Solve(input io.Reader) (answer string, err error) {
	var acc, lnum int
	var maskgrp uint64
	const grpsize int = 3
	lineIter := func(line string) {
		if 0 == lnum%grpsize {
			maskgrp = 0xffffffffffffffff
		}
		maskgrp &= ruckItemMask(line)
		if grpsize-1 == lnum%grpsize {
			acc += ruckPrioScan(maskgrp)
		}
		lnum++
	}

	puzzle.ForEachLine(input, lineIter)
	answer = strconv.Itoa(acc)
	return
}
