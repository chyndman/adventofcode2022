package aoc2022

import (
	"bufio"
	"io"
	"strconv"
)

type d13PacketNode struct {
	Next  *d13PacketNode
	IsNum bool
	Num   int
	List  *d13PacketNode
}

func d13PacketDecode(str string) (pkt *d13PacketNode) {
	pkt = &d13PacketNode{}
	path := make([]*d13PacketNode, 1)
	path[0] = pkt
	for i := 0; i < len(str); {
		switch str[i] {
		case '[':
			listpkt := &d13PacketNode{}
			path[len(path)-1].List = listpkt
			path = append(path, listpkt)
			i++
		case ']':
			path = path[0 : len(path)-1]
			i++
		case ',':
			next := &d13PacketNode{}
			path[len(path)-1].Next = next
			path[len(path)-1] = next
			i++
		default:
			for ; '0' <= str[i] && str[i] <= '9'; i++ {
				n := &(path[len(path)-1].Num)
				*n = *n*10 + int(str[i]-'0')
			}
			path[len(path)-1].IsNum = true
		}
	}
	return
}

func d13PacketsCompare(l, r *d13PacketNode) int {
	if l == nil && r == nil {
		return 0
	} else if r == nil {
		return -1
	} else if l == nil {
		return 1
	}

	lp, rp := l.Next, r.Next

	if l.IsNum && r.IsNum {
		if l.Num < r.Num {
			return 1
		} else if l.Num > r.Num {
			return -1
		}
	} else if !l.IsNum && !r.IsNum {
		if cmp := d13PacketsCompare(l.List, r.List); cmp != 0 {
			return cmp
		}
	} else {
		lp, rp = l, r
		cv := l
		if r.IsNum {
			cv = r
		}
		cv.IsNum = false
		cv.List = &d13PacketNode{}
		cv.List.IsNum = true
		cv.List.Num = cv.Num
	}

	return d13PacketsCompare(lp, rp)
}

type Day13Part1 struct{}

func (_ Day13Part1) Solve(input io.Reader) (answer string, err error) {
	acc := 0
	for pairIdx, sc := 1, bufio.NewScanner(input); ; pairIdx++ {
		sc.Scan()
		strl := sc.Text()
		l := d13PacketDecode(strl)

		sc.Scan()
		strr := sc.Text()
		r := d13PacketDecode(strr)

		cmp := d13PacketsCompare(l, r)
		if 1 == cmp {
			acc += pairIdx
		}

		if !sc.Scan() {
			break
		}
	}
	return strconv.Itoa(acc), nil
}

type Day13Part2 struct{}

func (_ Day13Part2) Solve(input io.Reader) (answer string, err error) {
	dla := d13PacketDecode("[[2]]")
	dlb := d13PacketDecode("[[6]]")
	idxa, idxb := 1, 2
	for sc := bufio.NewScanner(input); sc.Scan(); {
		if len(sc.Text()) == 0 {
			continue
		}
		p := d13PacketDecode(sc.Text())

		if 1 == d13PacketsCompare(p, dla) {
			idxa++
		}
		if 1 == d13PacketsCompare(p, dlb) {
			idxb++
		}
	}
	return strconv.Itoa(idxa * idxb), nil
}
