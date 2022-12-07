package aoc2022

import (
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"io"
	"regexp"
	"strconv"
)

type D7File struct {
	Size int
}

type D7Dir struct {
	Parent          *D7Dir
	Dirs            map[string]*D7Dir
	Files           map[string]D7File
	FileSizeShallow int
	FileSizeDeep    int
}

func newD7Dir(parent *D7Dir) *D7Dir {
	return &D7Dir{
		parent,
		make(map[string]*D7Dir),
		make(map[string]D7File),
		0, 0,
	}
}

var (
	cmdCdPat     = regexp.MustCompile(`\$ cd (.+)`)
	lsOutDirPat  = regexp.MustCompile(`dir (.+)`)
	lsOutFilePat = regexp.MustCompile(`(\d+) (.+)`)
)

type Day7 struct{}

var (
	Day7Part1 = Day7{}
)

func (p Day7) Solve(input io.Reader) (answer string, err error) {
	fs := newD7Dir(nil)
	current := fs

	lineIter := func(line string) {
		if cmdCd := cmdCdPat.FindStringSubmatch(line); nil != cmdCd {
			dest := cmdCd[1]
			switch dest {
			case "..":
				current = current.Parent
			case "/":
				current = fs
			default:
				current = current.Dirs[dest]
			}
		} else if lsOutDir := lsOutDirPat.FindStringSubmatch(line); nil != lsOutDir {
			name := lsOutDir[1]
			if _, exists := current.Dirs[name]; !exists {
				current.Dirs[name] = newD7Dir(current)
			}
		} else if lsOutFile := lsOutFilePat.FindStringSubmatch(line); nil != lsOutFile {
			size, _ := strconv.Atoi(lsOutFile[1])
			name := lsOutFile[2]
			sizeDiff := 0
			if _, exists := current.Files[name]; exists {
				sizeDiff -= size
			}
			current.Files[name] = D7File{size}
			sizeDiff += size

			current.FileSizeShallow += sizeDiff
			for d := current; nil != d; d = d.Parent {
				d.FileSizeDeep += sizeDiff
			}
		}
	}

	puzzle.ForEachLine(input, lineIter)

	acc := 0
	dl := make([]*D7Dir, 1)
	dl[0] = fs
	for 0 < len(dl) {
		head := dl[0]
		dl = dl[1:]

		if head.FileSizeDeep <= 100000 {
			acc += head.FileSizeDeep
		}

		for _, subdir := range head.Dirs {
			dl = append(dl, subdir)
		}
	}

	answer = strconv.Itoa(acc)
	return
}
