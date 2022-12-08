package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"strconv"
)

type D8Tree struct {
	Height uint
}

type Day8 struct {
	SolveWithTrees func(trees *[][]D8Tree) int
}

func swt1(trees *[][]D8Tree) (acc int) {
	dxs := [4]int{1, 0, -1, 0}
	dys := [4]int{0, -1, 0, 1}
	for y := 0; y < len(*trees); y++ {
		for x := 0; x < len((*trees)[y]); x++ {
			visible := false
			for di := 0; !visible && di < 4; di++ {
				xp, yp := x, y
				for {
					xp += dxs[di]
					yp += dys[di]
					if xp < 0 || len((*trees)[y]) <= xp || yp < 0 || len(*trees) <= yp {
						visible = true
						break
					} else if (*trees)[y][x].Height <= (*trees)[yp][xp].Height {
						break
					}
				}
			}
			if visible {
				acc++
			}
		}
	}
	return
}

func swt2(trees *[][]D8Tree) (scoreMax int) {
	dxs := [4]int{1, 0, -1, 0}
	dys := [4]int{0, -1, 0, 1}
	for y := 1; y < len(*trees)-1; y++ {
		for x := 1; x < len((*trees)[y])-1; x++ {
			score := 1
			for di := 0; di < 4; di++ {
				dist := 0
				xp, yp := x, y
				for {
					xp += dxs[di]
					yp += dys[di]
					if xp < 0 || len((*trees)[y]) <= xp || yp < 0 || len(*trees) <= yp {
						break
					}
					dist++
					if (*trees)[y][x].Height <= (*trees)[yp][xp].Height {
						break
					}
				}
				score *= dist
			}
			if scoreMax < score {
				scoreMax = score
			}
		}
	}
	return
}

var (
	Day8Part1 = Day8{swt1}
	Day8Part2 = Day8{swt2}
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
	answer = strconv.Itoa(p.SolveWithTrees(&trees))
	return
}
