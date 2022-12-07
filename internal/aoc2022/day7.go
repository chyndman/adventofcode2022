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

type Day7 struct{
	SolveWithFs func(*D7Dir) (string, error)
}

func bfsTraverseD7Dir(fs *D7Dir, fn func(dir *D7Dir)) {
	dl := make([]*D7Dir, 1)
	dl[0] = fs
	for 0 < len(dl) {
		head := dl[0]
		dl = dl[1:]

		fn(head)

		for _, subdir := range head.Dirs {
			dl = append(dl, subdir)
		}
	}
}

func swfs1(fs *D7Dir) (answer string, err error) {
	acc := 0
	dirFn := func(dir *D7Dir) {
		if dir.FileSizeDeep <= 100000 {
			acc += dir.FileSizeDeep
		}		
	}
	bfsTraverseD7Dir(fs, dirFn)
	answer = strconv.Itoa(acc)
	return
}

func swfs2(fs *D7Dir) (answer string, err error) {
	best := fs.FileSizeDeep
	reclaimMin := fs.FileSizeDeep - (70000000 - 30000000)
	if 0 >= reclaimMin {
		return "0", nil
	}

	dirFn := func(dir *D7Dir) {
		if dir.FileSizeDeep >= reclaimMin && best > dir.FileSizeDeep {
			best = dir.FileSizeDeep
		}
	}
	bfsTraverseD7Dir(fs, dirFn)
	answer = strconv.Itoa(best)
	return
}

var (
	Day7Part1 = Day7{swfs1}
	Day7Part2 = Day7{swfs2}
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
	return p.SolveWithFs(fs)
}
