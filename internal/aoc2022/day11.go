package aoc2022

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	d11StartingPat = regexp.MustCompile("  Starting items: (.+)")
	d11OpPat       = regexp.MustCompile(`  Operation: new = old (\*|\+) (.+)`)
	d11TestPat     = regexp.MustCompile(`  Test: divisible by (\d+)`)
	d11ThrowPat    = regexp.MustCompile(`    If (true|false): throw to monkey (\d+)`)
)

type d11Monkey struct {
	Items                           []uint64
	Op                              func(uint64) uint64
	ThrowMod, TrueThrow, FalseThrow uint64
	InspectionCount                 uint64
}

type Day11 struct {
	Rounds, WorryDiv uint64
}

var (
	Day11Part1 = Day11{20, 3}
	Day11Part2 = Day11{10000, 1}
)

func (p Day11) Solve(input io.Reader) (answer string, err error) {
	var acc uint64
	monkeys := make([]d11Monkey, 0)
	var lcm uint64 = uint64(1)
	for sc := bufio.NewScanner(input); sc.Scan(); {
		mk := d11Monkey{}

		sc.Scan()
		itemStrs := strings.Split(d11StartingPat.FindStringSubmatch(sc.Text())[1], ", ")
		mk.Items = make([]uint64, len(itemStrs))
		for i := range mk.Items {
			item, _ := strconv.Atoi(itemStrs[i])
			mk.Items[i] = uint64(item)
		}

		sc.Scan()
		opSubmatch := d11OpPat.FindStringSubmatch(sc.Text())
		operator, operand := opSubmatch[1], opSubmatch[2]
		binop := func(a, b uint64) uint64 { return a + b }
		if '*' == operator[0] {
			binop = func(a, b uint64) uint64 { return a * b }
		}
		op := func(x uint64) uint64 { return binop(x, x) }
		if "old" != operand {
			num, _ := strconv.Atoi(operand)
			op = func(x uint64) uint64 { return binop(x, uint64(num)) }
		}
		mk.Op = op

		var n int

		sc.Scan()
		n, _ = strconv.Atoi(d11TestPat.FindStringSubmatch(sc.Text())[1])
		mk.ThrowMod = uint64(n)
		lcm *= mk.ThrowMod

		sc.Scan()
		n, _ = strconv.Atoi(d11ThrowPat.FindStringSubmatch(sc.Text())[2])
		mk.TrueThrow = uint64(n)

		sc.Scan()
		n, _ = strconv.Atoi(d11ThrowPat.FindStringSubmatch(sc.Text())[2])
		mk.FalseThrow = uint64(n)

		sc.Scan() // empty line
		monkeys = append(monkeys, mk)
	}

	for round := uint64(0); round < p.Rounds; round++ {
		for i := range monkeys {
			for _, itemWorry := range monkeys[i].Items {
				monkeys[i].InspectionCount++
				newItemWorry := monkeys[i].Op(itemWorry) / p.WorryDiv
				newItemWorry = newItemWorry % lcm
				catchIdx := monkeys[i].FalseThrow
				if 0 == newItemWorry%monkeys[i].ThrowMod {
					catchIdx = monkeys[i].TrueThrow
				}
				monkeys[catchIdx].Items = append(monkeys[catchIdx].Items, newItemWorry)
			}
			monkeys[i].Items = monkeys[i].Items[0:0]
		}
	}

	for i := range monkeys {
		for j := i + 1; j < len(monkeys); j++ {
			monkeyBusiness := monkeys[i].InspectionCount * monkeys[j].InspectionCount
			if acc < monkeyBusiness {
				acc = monkeyBusiness
			}
		}
	}
	return strconv.FormatUint(acc, 10), nil
}
