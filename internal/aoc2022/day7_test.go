package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzleexpect"
	"testing"
)

func TestDay7(t *testing.T) {
	const vec string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	t.Run("1", puzzleexpect.SolveTester(Day7Part1, vec, "95437"))
}
