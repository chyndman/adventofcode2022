package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

const (
	east  = 1 << 0
	north = 1 << 1
	west  = 1 << 2
	south = 1 << 3
)

type D8Tree struct {
	Height uint
}

type Day8 struct{}

var (
	Day8Part1 = Day8{}
)

func (p Day8) Solve(input io.Reader) (answer string, err error) {
	var trees [][]D8Tree
	var width, height int
	lineIter := func(line string) {
		if nil == trees {
			width = len(line)
			trees = make([][]D8Tree, 0, width)
		}

		trees = trees[:height+1]
		trees[height] = make([]D8Tree, width)

		for x := 0; x < width; x++ {
			trees[height][x] = D8Tree{uint(line[x] - '0')}
		}

		height++
	}

	puzzle.ForEachLine(input, lineIter)

	acc := 0
	dxs := [4]int{0, 1, 0, -1}
	dys := [4]int{1, 0, -1, 0}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			visible := false
			for di := 0; !visible && di < 4; di++ {
				xp, yp := x, y
				for {
					xp += dxs[di]
					yp += dys[di]
					if xp < 0 || width <= xp || yp < 0 || height <= yp {
						visible = true
						break
					} else if trees[y][x].Height <= trees[yp][xp].Height {
						break
					}
				}
			}
			if visible {
				acc++
			}
		}
	}

	answer = strconv.Itoa(acc)
	return
}
