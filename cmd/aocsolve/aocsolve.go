package main

import (
	"fmt"
	"github.com/chyndman/adventofcode2022/internal/aoc2022"
	"github.com/chyndman/adventofcode2022/internal/demo"
	"github.com/chyndman/adventofcode2022/internal/puzzle"
	"log"
	"os"
)

var PuzzleSets = map[string]map[string]puzzle.Puzzle{
	"demo":    demo.Puzzles,
	"aoc2022": aoc2022.Puzzles,
}

func getArg(index int, valDefault string) (val string) {
	if index < len(os.Args) {
		val = os.Args[index]
	} else {
		val = valDefault
	}
	return
}

func main() {
	param := map[string]string{
		"pzset":  getArg(1, ""),
		"pz":     getArg(2, ""),
		"infile": getArg(3, "input.txt"),
	}

	if "" == param["pzset"] || "" == param["pz"] {
		log.Fatal("Usage: aocsolve <puzzle set> <puzzle> [input file]")
	}

	pzSet := PuzzleSets[param["pzset"]]
	if nil == pzSet {
		log.Fatalf("No puzzle set %s", param["pzset"])
	}

	pz := pzSet[param["pz"]]
	if nil == pz {
		log.Fatalf("No puzzle in set %s with name %s", param["pzset"], param["pz"])
	}

	f, err := os.Open(param["infile"])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Solving %s:%s for file %s", param["pzset"], param["pz"], param["infile"])
	pzAnswer, pzErr := pz.Solve(f)
	if pzErr != nil {
		log.Print(pzErr)
	} else {
		fmt.Println(pzAnswer)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
