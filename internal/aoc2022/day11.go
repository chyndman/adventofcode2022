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
	Items                           []int
	Op                              func(int) int
	ThrowMod, TrueThrow, FalseThrow int
	InspectionCount                 int
}

type Day11 struct{}

var (
	Day11Part1 = Day11{}
)

func (p Day11) Solve(input io.Reader) (answer string, err error) {
	acc := 0
	monkeys := make([]d11Monkey, 0)
	for sc := bufio.NewScanner(input); sc.Scan(); {
		mk := d11Monkey{}

		sc.Scan()
		itemStrs := strings.Split(d11StartingPat.FindStringSubmatch(sc.Text())[1], ", ")
		mk.Items = make([]int, len(itemStrs))
		for i := range mk.Items {
			mk.Items[i], _ = strconv.Atoi(itemStrs[i])
		}

		sc.Scan()
		opSubmatch := d11OpPat.FindStringSubmatch(sc.Text())
		operator, operand := opSubmatch[1], opSubmatch[2]
		binop := func(a, b int) int { return a + b }
		if '*' == operator[0] {
			binop = func(a, b int) int { return a * b }
		}
		op := func(x int) int { return binop(x, x) }
		if "old" != operand {
			num, _ := strconv.Atoi(operand)
			op = func(x int) int { return binop(x, num) }
		}
		mk.Op = op

		sc.Scan()
		mk.ThrowMod, _ = strconv.Atoi(d11TestPat.FindStringSubmatch(sc.Text())[1])

		sc.Scan()
		mk.TrueThrow, _ = strconv.Atoi(d11ThrowPat.FindStringSubmatch(sc.Text())[2])

		sc.Scan()
		mk.FalseThrow, _ = strconv.Atoi(d11ThrowPat.FindStringSubmatch(sc.Text())[2])

		sc.Scan() // empty line
		monkeys = append(monkeys, mk)
	}

	for round := 0; round < 20; round++ {
		for i := range monkeys {
			for _, itemWorry := range monkeys[i].Items {
				monkeys[i].InspectionCount++
				newItemWorry := monkeys[i].Op(itemWorry) / 3
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
	return strconv.Itoa(acc), nil
}
