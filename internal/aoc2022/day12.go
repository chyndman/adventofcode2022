package aoc2022

import (
	"container/heap"
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"image"
	"io"
	"math"
	"strconv"
)

type d12Heightmap []string

func (hm d12Heightmap) Find(c byte) (pt image.Point) {
	for y := range hm {
		for x := range hm[y] {
			if c == hm[y][x] {
				pt.X, pt.Y = x, y
				break
			}
		}
	}
	return
}

func (hm d12Heightmap) HeightAt(pt image.Point) int {
	if 0 <= pt.Y && pt.Y < len(hm) && 0 <= pt.X && pt.X < len(hm[pt.Y]) {
		c := hm[pt.Y][pt.X]
		switch c {
		case 'S':
			c = 'a'
		case 'E':
			c = 'z'
		}
		return int(c - 'a')
	}
	return math.MinInt
}

func (hm d12Heightmap) NeighborsAt(pt image.Point) (npts []image.Point) {
	npts = make([]image.Point, 0, 4)
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	h := hm.HeightAt(pt)
	for i := 0; i < 4; i++ {
		npt := image.Point{pt.X + dx[i], pt.Y + dy[i]}
		hp := hm.HeightAt(npt)
		dh := hp - h
		if hp >= 0 && dh <= 1 {
			npts = append(npts, npt)
		}
	}

	return
}

type d12PathfindVertex struct {
	Prio uint
	Loc  image.Point
}

type d12PathfindHeap []d12PathfindVertex

func (h d12PathfindHeap) Len() int {
	return len(h)
}

func (h d12PathfindHeap) Less(i, j int) bool {
	return h[i].Prio < h[j].Prio
}

func (h d12PathfindHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *d12PathfindHeap) Push(x any) {
	*h = append(*h, x.(d12PathfindVertex))
}

func (h *d12PathfindHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (hm d12Heightmap) Pathfind(source image.Point) (dists [][]uint) {
	dists = make([][]uint, len(hm))
	for y := range dists {
		dists[y] = make([]uint, len(hm[y]))
		for x := range dists[y] {
			dists[y][x] = math.MaxUint
		}
	}

	h := make(d12PathfindHeap, 1)
	h[0].Loc = source
	heap.Init(&h)

	for 0 < h.Len() {
		u := heap.Pop(&h).(d12PathfindVertex)
		if u.Prio < dists[u.Loc.Y][u.Loc.X] {
			dists[u.Loc.Y][u.Loc.X] = u.Prio
		}
		stepdist := dists[u.Loc.Y][u.Loc.X] + 1
		for _, v := range hm.NeighborsAt(u.Loc) {
			if stepdist < dists[v.Y][v.X] {
				pv := d12PathfindVertex{
					Prio: stepdist,
					Loc:  v,
				}
				heap.Push(&h, pv)
				dists[v.Y][v.X] = stepdist
			}
		}
	}

	return
}

func swhm1(hm d12Heightmap) uint {
	dists := hm.Pathfind(hm.Find('S'))
	dest := hm.Find('E')
	return dists[dest.Y][dest.X]
}

func swhm2(hm d12Heightmap) (distMin uint) {
	distMin = math.MaxUint
	dest := hm.Find('E')
	for y := range hm {
		for x := range hm[y] {
			if 'a' == hm[y][x] || 'S' == hm[y][x] {
				dists := hm.Pathfind(image.Point{x, y})
				dist := dists[dest.Y][dest.X]
				if dist < distMin {
					distMin = dist
				}
			}
		}
	}
	return
}

type Day12 struct {
	SolveWithHeightmap func(hm d12Heightmap) uint
}

var (
	Day12Part1 = Day12{swhm1}
	Day12Part2 = Day12{swhm2}
)

func (p Day12) Solve(input io.Reader) (answer string, err error) {
	hm := d12Heightmap(puzzle.ReadLines(input))
	return strconv.Itoa(int(p.SolveWithHeightmap(hm))), nil
}
